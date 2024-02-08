package xenapi

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
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
		return nil, fmt.Errorf("Could not create request for %v() : %w", request.Method, err)
	}

	httpResponse, err := client.httpClient.Do(httpRequest)
	fmt.Println("--------------httpResponse")
	fmt.Println(httpResponse)

	if err != nil {
		return nil, fmt.Errorf("Call %v() on %v. Error making http request: %w", request.Method, httpRequest.URL.Redacted(), err)
	}
	defer httpResponse.Body.Close()

	var rpcResponse *Response
	// decoder := json.NewDecoder(httpResponse.Body)
	// decoder.DisallowUnknownFields()
	// decoder.UseNumber()
	// err = decoder.Decode(&rpcResponse)

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("Call %v() on %v status code: %v. Could not read response body: %w", request.Method, httpRequest.URL.Redacted(), httpResponse.StatusCode, err)
	}
	body = ConvertJsonValueWhenNaNOrInf(body)
	err = json.Unmarshal(body, &rpcResponse)
	// parsing error
	if err != nil && err != io.EOF {
		// fmt.Printf("error decoding sakura response: %v", err)
		// fmt.Printf("sakura response: %q", body)
		return nil, fmt.Errorf("Call %v() on %v status code: %v. Could not decode response body: %w", request.Method, httpRequest.URL.Redacted(), httpResponse.StatusCode, err)
	}

	// response body empty
	if rpcResponse == nil {
		return nil, fmt.Errorf("Call %v() on %v status code: %v. Response missing", request.Method, httpRequest.URL.Redacted(), httpResponse.StatusCode)
	}

	return rpcResponse, nil
}

type ClientOpts struct {
	Url        string
	CaCertPath string
}

func NewJsonRPCClient(opts *ClientOpts) *JsonRpcClient {
	httpClient := &http.Client{}
	if strings.HasPrefix(opts.Url, "https://") {
		caCertPool := x509.NewCertPool()
		if opts.CaCertPath != "" {
			caCert, err := os.ReadFile(opts.CaCertPath)
			if err != nil {
				log.Fatal(err)
			}
			caCertPool.AppendCertsFromPEM(caCert)
		}
		tlsConfig := &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: false,
			RootCAs:            caCertPool,
		}
		transport := &http.Transport{
			TLSClientConfig: tlsConfig,
			IdleConnTimeout: 30 * time.Second,
		}
		httpClient.Transport = transport
	}

	// set up jsonrpc client
	client := &JsonRpcClient{
		endpoint:   fmt.Sprintf("%s%s", opts.Url, "/jsonrpc"),
		httpClient: httpClient,
	}

	return client
}
