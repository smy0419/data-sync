package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/service"
)

type ContractTemplateEvent struct{}

var contractTemplateService = service.ContractTemplateService{}

func (contractTemplateEvent ContractTemplateEvent) handle(b blockInfo, args map[string]interface{}) error {
	category, ok := args["category"]
	if !ok {
		return contractTemplateEvent.checkArg("category")
	}

	name, ok := args["name"]
	if !ok {
		return contractTemplateEvent.checkArg("name")
	}

	owner, ok := args["owner"]
	if !ok {
		return contractTemplateEvent.checkArg("owner")
	}

	key, ok := args["key"]
	if !ok {
		return contractTemplateEvent.checkArg("key")
	}

	approvers, ok := args["approvers"]
	if !ok {
		return contractTemplateEvent.checkArg("approvers")
	}

	rejecters, ok := args["rejecters"]
	if !ok {
		return contractTemplateEvent.checkArg("rejecters")
	}

	allApprover, ok := args["allApprover"]
	if !ok {
		return contractTemplateEvent.checkArg("allApprover")
	}

	status, ok := args["status"]
	if !ok {
		return contractTemplateEvent.checkArg("status")
	}

	return contractTemplateService.Insert(
		b.height,
		b.time,
		owner.(*asimovCommon.Address).Hex(),
		*category.(*uint16),
		*name.(*string),
		asimovCommon.Bytes2Hex((*key.(*[32]uint8))[:]),
		*approvers.(*uint8),
		*rejecters.(*uint8),
		*allApprover.(*uint8),
		*status.(*uint8),
	)
}

func (contractTemplateEvent ContractTemplateEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("contract template event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
