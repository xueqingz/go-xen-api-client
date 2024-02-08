package xenapi

import (
	"context"
	"fmt"
)

func (client *Client) APICall(method string, params ...interface{}) (result interface{}, err error) {
	fmt.Println("\nmethod:" + method)
	fmt.Println(params...)

	response, err := client.rpc.Call(context.Background(), method, params...)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println("\nresponse:")
	// fmt.Println(response.Result)
	// fmt.Println(response.Error)

	if response.Error != nil {
		fmt.Println("Continue response.Error:")
		errString := fmt.Sprintf("API Error: code %d, message %s", response.Error.Code, response.Error.Message)
		if response.Error.Data != nil {
			errString = errString + fmt.Sprintf(", data %v", response.Error.Data)
		}
		err = fmt.Errorf(errString)
		fmt.Println(err.Error())
		return
	}
	// fmt.Println("The response result is:")
	// fmt.Println(response.Result)

	result = response.Result
	return
}

func NewClient(opts *ClientOpts) (*Client, error) {
	rpc := NewJsonRPCClient(opts)

	return prepClient(rpc), nil
}
