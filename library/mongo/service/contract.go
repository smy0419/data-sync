package service

import (
	"errors"
	"github.com/AsimovNetwork/asimov/rpcs/rpcjson"
	"github.com/AsimovNetwork/data-sync/library/common"
)

type ContractService struct{}

func (contractService ContractService) GetTemplate(address string) (*rpcjson.ContractTemplate, error) {
	param := common.NewChainRequest("getContractTemplate", []interface{}{address})
	result, ok := common.Post(common.Cfg.BlockChainRpc, param)
	if !ok {
		return nil, errors.New("call block chain failed")
	}
	resultMap := (result).(map[string]interface{})

	contractTemplate := &rpcjson.ContractTemplate{}
	err := common.ToStruct(resultMap, contractTemplate)
	// if contractTemplate.TemplateType == 0 || contractTemplate.TemplateTName == "" {
	// 	return nil, errors.New(fmt.Sprintf("contract %s template is not exist", address))
	// }
	return contractTemplate, err
}
