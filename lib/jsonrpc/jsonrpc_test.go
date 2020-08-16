package jsonrpc

import (
	"fmt"
	"github.com/GincoInc/jsonrpc"
	"testing"
	"time"
)

// RPCClient ...
type RPCClient struct {
	RPCClient *jsonrpc.RPCClient
}

func NewRPCClient(endpoint string) *RPCClient {
	return &RPCClient{
		RPCClient: jsonrpc.NewRPCClientWithTimeout(endpoint, 1 * time.Second),
	}
}

type Params struct {
	Account 		string 	`json:"account"`
	Strict 			bool	`json:"strict"`
	LedgerIndex		string	`json:"ledger_index"`
	Queue			bool	`json:"queue"`
}

func RunRPC(){


	//resp, err:=c.RPCClient.Call("ledger_current", []interface{}{})
	//if err != nil {
	//	return
	//}


	//params := Params{
	//	Account: "rGJx52axFJ4VhBqviwhzMFibC6dTwR94ww",
	//	Strict: true,
	//	LedgerIndex: "current",
	//	Queue: true,
	//}
	//resp, err := c.RPCClient.Call("account_info", []Params{params})
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}



}

func Test_jsonrpc(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "timeout",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewRPCClient("http://29.23.12.41:5005")
			resp, err:= c.RPCClient.Call("account_info", []interface{}{
				map[string]interface{}{
					"account"		: "rGJx52axFJ4VhBqviwhzMFibC6dTwR94ww",
					"strict"		: true,
					"ledger_index"	: "current",
					"queue"			: true,
				},
			})
			if err != nil {
				fmt.Println("err: ", err)
				return
			}

			fmt.Printf("res: %v", resp)
		})
	}
}
