package common

import (
	"errors"
	"github.com/AsimovNetwork/asimov/rpcs/rpcjson"
)

func GetContractTemplateInfoByName(category uint16, name string) (*rpcjson.ContractTemplateDetail, error) {
	param := NewChainRequest("getContractTemplateInfoByName", []interface{}{category, name})
	result, ok := Post(Cfg.BlockChainRpc, param)
	if !ok {
		return nil, errors.New("call block chain failed")
	}
	resultMap := (result).(map[string]interface{})

	bestBlock := &rpcjson.ContractTemplateDetail{}
	err := ToStruct(resultMap, bestBlock)
	return bestBlock, err
}
