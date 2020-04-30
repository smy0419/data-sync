package service

import (
	"errors"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/asimov/vm/fvm"
	"github.com/AsimovNetwork/data-sync/library/common"
)

type SystemContractService struct{}

func (systemContractService SystemContractService) GetSystemContractAbi(height int64, address string) (bool, string, error) {
	// param["params"] = []interface{}{int32(height), vm.RevertSystemContractCode(asimovCommon.HexToAddress(address[2:]))}
	param := common.NewChainRequest("getGenesisContractByHeight", []interface{}{int32(height), address})
	result, ok := common.Post(common.Cfg.BlockChainRpc, param)
	if !ok {
		return false, "", errors.New("get system contract abi failed")
	}
	resultMap := (result).(map[string]interface{})
	exist, ok := resultMap["exist"]
	if !ok {
		return false, "", errors.New("get system contract abi failed")
	}
	abi, ok := resultMap["abi"]
	if !ok {
		return false, "", errors.New("get system contract abi failed")
	}

	return exist.(bool), abi.(string), nil
}

func (systemContractService SystemContractService) GetDaoContractAbiByAddress(address string) (bool, string, error) {
	contractTemplate, err := contractService.GetTemplate(address)
	if err != nil {
		return false, "", errors.New("get dao contract abi failed")
	}

	if int(contractTemplate.TemplateType) == common.Category && contractTemplate.TemplateTName == common.TemplateDAO {
		return true, common.DaoABI, nil
	} else if int(contractTemplate.TemplateType) == common.Category && contractTemplate.TemplateTName == common.TemplateVote {
		return true, common.VoteABI, nil
	} else {
		return false, "", nil
	}
}

func (systemContractService SystemContractService) GetContractTemplateInfoAddress(address string) (interface{}, error) {
	funcName := "getTemplateInfo"
	data, err := fvm.PackFunctionArgs(common.TemplateABI, funcName)
	if err != nil {
		return nil, err
	}

	param := common.NewChainRequest("callReadOnlyFunction", []interface{}{common.OfficialCaller, address, asimovCommon.Bytes2Hex(data), funcName, common.TemplateABI})
	result, ok := common.Post(common.Cfg.BlockChainRpc, param)
	if !ok {
		return nil, errors.New("call block chain failed")
	}
	return result, nil
}
