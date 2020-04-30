package common

import (
	"errors"
)

func SendRawTransaction(params []interface{}) (string, error) {
	param := NewChainRequest("sendRawTransaction", params)
	result, success := Post(Cfg.BlockChainRpc, param)
	if !success {
		return "", errors.New("rpc:sendRawTransaction failed")
	}
	return result.(string), nil
}
