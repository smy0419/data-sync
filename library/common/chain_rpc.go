package common

import (
	"fmt"
)

const RPCPrefix = "asimov_"

type ChainRequest struct {
	ID      int64         `json:"id"`
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

func NewChainRequest(method string, params []interface{}) ChainRequest {
	return ChainRequest{
		ID:      Now(),
		JsonRpc: "2.0",
		Method:  fmt.Sprintf("%s%s", RPCPrefix, method),
		Params:  params,
	}
}
