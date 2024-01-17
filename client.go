package xenapi

import (
	"context"
	"fmt"
	"strconv"
)

func (client *Client) APICall(method string, params ...interface{}) (result interface{}, err error) {
	fmt.Println("\nmethod:" + method)
	fmt.Println(params...)

	response, err := client.rpc.Call(context.Background(), method, params...)
	if err != nil {
		fmt.Println("\nhttp error:")
		fmt.Println(err.Error())
		return
	}

	if response.Error != nil {
		fmt.Println("Continue response.Error:")
		fmt.Println(response.Error)
		err = fmt.Errorf("Error code: " + strconv.Itoa(response.Error.Code) + ", message: " + response.Error.Message)
		return
	}
	// fmt.Println("The response result is:")
	// fmt.Println(response.Result)

	result = response.Result
	return
}

func NewClient(url string) (*Client, error) {
	rpc := NewJsonRPCClient(fmt.Sprintf("%s%s", url, "/jsonrpc"))

	return prepClient(rpc), nil

}
