package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/service"
)

type RenameOrganizationEvent struct{}

var daoOrganizationService = service.DaoOrganizationService{}

func (renameOrganizationEvent RenameOrganizationEvent) handle(b blockInfo, args map[string]interface{}) error {
	contractAddress, ok := args["contractAddress"]
	if !ok {
		return renameOrganizationEvent.checkArg("contractAddress")
	}

	newName, ok := args["newName"]
	if !ok {
		return renameOrganizationEvent.checkArg("newName")
	}

	err := daoOrganizationService.UpdateOrgName(b.height, contractAddress.(*asimovCommon.Address).Hex(), *newName.(*string))
	if err != nil {
		return err
	}

	err = mysqlDaoOrganizationService.UpdateByContractAddress(contractAddress.(*asimovCommon.Address).Hex(), *newName.(*string))
	if err != nil {
		return err
	}

	return nil
}

func (renameOrganizationEvent RenameOrganizationEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("dao rename organization event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
