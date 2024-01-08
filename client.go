package xenapi

import (
	"fmt"
	"context"
	"github.com/ybbus/jsonrpc/v3"
)


func (client *Client) APICall(method string, params ...interface{}) (result interface{}, err error) {
	fmt.Println(method)
	fmt.Println(params...)

	response, err := client.rpc.Call(context.Background(), method, params...)
	if err != nil {
		fmt.Println("http error:")
		fmt.Println(err)
		return
	}
	
	if response.Error != nil {
		fmt.Println("Continue response.Error:")
		fmt.Println(response.Error)
		err = fmt.Errorf("%d:%s",response.Error.Code, response.Error.Message)
		return
    }
	// fmt.Println("The response result is:")
	// fmt.Println(response.Result)

	result = response.Result
	return
}

func NewClient(url string) (*Client, error) {
	// rpc := jsonrpc.NewClientWithOpts(fmt.Sprintf("%s%s", url,"/jsonrpc"), &jsonrpc.RPCClientOpts{
	// 	CustomHeaders: map[string]string{
	// 		"Content-Type": "application/json",
	// 	},
	// })
	rpc := jsonrpc.NewClient(fmt.Sprintf("%s%s", url,"/jsonrpc"))

	return prepClient(rpc), nil

}