package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	mongodb "github.com/AsimovNetwork/data-sync/library/mongo/service"
	"github.com/AsimovNetwork/data-sync/library/mysql/service"
)

type CloseOrganizationEvent struct{}

var contractService = mongodb.ContractService{}

var daoMessageService = service.DaoMessageService{}

func (closeOrganizationEvent CloseOrganizationEvent) handle(b blockInfo, args map[string]interface{}) error {
	contractAddress, ok := args["contractAddress"]
	if !ok {
		return closeOrganizationEvent.checkArg("contractAddress")
	}

	contractTemplate, err := mongodb.ContractService{}.GetTemplate(contractAddress.(*asimovCommon.Address).Hex())
	if err != nil {
		return err
	}

	if int(contractTemplate.TemplateType) != common.Category || contractTemplate.TemplateTName != common.TemplateDAO {
		return nil
	}

	// update dao_organization of mongodb
	err = daoOrganizationService.CloseOrg(b.height, contractAddress.(*asimovCommon.Address).Hex())
	if err != nil {
		return err
	}

	// update t_dao_organization of mysql
	err = mysqlDaoOrganizationService.CloseOrg(contractAddress.(*asimovCommon.Address).Hex(), b.height)
	if err != nil {
		return err
	}
	return nil
}

func (closeOrganizationEvent CloseOrganizationEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("dao close organization event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
