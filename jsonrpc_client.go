package xenapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
)

type Request struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      int         `json:"id"`
}

type ResponseError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Response struct {
	JSONRPC string         `json:"jsonrpc"`
	Result  interface{}    `json:"result,omitempty"`
	Error   *ResponseError `json:"error,omitempty"`
	ID      int            `json:"id"`
}

type HTTPError struct {
	Code int
	err  error
}

func (e *HTTPError) Error() string {
	return e.err.Error()
}

func Params(params ...interface{}) interface{} {
	var finalParams interface{}
	finalParams = params
	if len(params) == 1 {
		if params[0] != nil {
			var typeOf reflect.Type
			for typeOf = reflect.TypeOf(params[0]); typeOf != nil && typeOf.Kind() == reflect.Ptr; typeOf = typeOf.Elem() {
			}
			typeArr := []reflect.Kind{reflect.Struct, reflect.Array, reflect.Slice, reflect.Interface, reflect.Map}
			if typeOf != nil {
				for _, value := range typeArr {
					if value == typeOf.Kind() {
						finalParams = params[0]
						break
					}
				}
			}
		}
	}
	return finalParams
}

type JsonRpcClient struct {
	endpoint   string
	httpClient *http.Client
}

func (client *JsonRpcClient) newRequest(ctx context.Context, req interface{}) (*http.Request, error) {
	dataByte, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", client.endpoint, bytes.NewReader(dataByte))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Accept", "application/json")

	return request, nil
}

func ConvertJsonValueWhenNaNOrInf(jsonBytes []byte) []byte {
	jsonStr1 := string(jsonBytes)
	jsonStr1 = strings.Replace(jsonStr1, ":Infinity", ":\"+Inf\"", -1)
	jsonStr1 = strings.Replace(jsonStr1, ":-Infinity", ":\"-Inf\"", -1)
	// 将NaN替换为null
	jsonStr1 = strings.Replace(jsonStr1, ":NaN,", ":null", -1)

	return []byte(jsonStr1)
}

func (client *JsonRpcClient) Call(ctx context.Context, method string, params ...interface{}) (*Response, error) {
	request := &Request{
		ID:      0,
		Method:  method,
		Params:  Params(params...),
		JSONRPC: "2.0",
	}

	httpRequest, err := client.newRequest(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("rpc call %v() on %v: %w", request.Method, client.endpoint, err)
	}

	httpResponse, err := client.httpClient.Do(httpRequest)
	fmt.Println("--------------httpResponse")
	fmt.Println(httpResponse)

	if err != nil {
		return nil, fmt.Errorf("rpc call %v() on %v: %w", request.Method, httpRequest.URL.Redacted(), err)
	}
	defer httpResponse.Body.Close()

	var rpcResponse *Response
	// decoder := json.NewDecoder(httpResponse.Body)
	// decoder.DisallowUnknownFields()
	// decoder.UseNumber()
	// err = decoder.Decode(&rpcResponse)

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("rpc call %v() on %v status code: %v. could not read from response body: %w", request.Method, httpRequest.URL.Redacted(), httpResponse.StatusCode, err)
	}
	body = ConvertJsonValueWhenNaNOrInf(body)
	err = json.Unmarshal(body, &rpcResponse)
	// parsing error
	if err != nil && err != io.EOF {
		fmt.Printf("error decoding sakura response: %v", err)
		// fmt.Printf("sakura response: %q", body)
		return nil, fmt.Errorf("rpc call %v() on %v status code: %v. could not decode body to rpc response: %w", request.Method, httpRequest.URL.Redacted(), httpResponse.StatusCode, err)
	}

	// response body empty
	if rpcResponse == nil {
		return nil, fmt.Errorf("rpc call %v() on %v status code: %v. rpc response missing", request.Method, httpRequest.URL.Redacted(), httpResponse.StatusCode)
	}

	return rpcResponse, nil
}

func NewJsonRPCClient(endpoint string) *JsonRpcClient {
	client := &JsonRpcClient{
		endpoint:   endpoint,
		httpClient: &http.Client{},
	}

	return client
}
