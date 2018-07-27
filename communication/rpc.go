package communication

import (
	"github.com/GincoInc/jsonrpc"
	"fmt"
)

// RPCClient ...
type RPCClient struct {
	RPCClient *jsonrpc.RPCClient
}

func NewRPCClient(endpoint string) *RPCClient {
	return &RPCClient{
		RPCClient: jsonrpc.NewRPCClient(endpoint),
	}
}

//type Params struct {
//
//}

type Params struct {
	Account 		string 	`json:"account"`
	Strict 			bool	`json:"strict"`
	LedgerIndex		string	`json:"ledger_index"`
	Queue			bool	`json:"queue"`
}

func RunRPC(){
	c := NewRPCClient("http://localhost:5005")

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
}