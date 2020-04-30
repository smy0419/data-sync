package service

import (
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"github.com/AsimovNetwork/data-sync/library/mysql"
	models "github.com/AsimovNetwork/data-sync/library/mysql/model"
)

type MysqlDaoOrganizationService struct{}

func (mysqlDaoOrganizationService MysqlDaoOrganizationService) CreateOrg(txHash string, contractAddress string, voteContractAddress string) error {
	now := (int)(common.NowSecond())

	sql := "update t_dao_organization set state = ?, contract_address = ?, vote_contract_address = ?, update_time = ? where tx_hash = ?"
	_, err := mysql.Engine.Exec(sql, model.OrgStatusNormal, contractAddress, voteContractAddress, now, txHash)

	if err != nil {
		common.Logger.Error(err)
		return err
	}
	return nil
}

func (mysqlDaoOrganizationService MysqlDaoOrganizationService) CloseOrg(contractAddress string, height int64) error {
	now := (int)(common.NowSecond())

	sql := "update t_dao_organization set state = ?, update_time = ? where contract_address = ?"
	_, err := mysql.Engine.Exec(sql, model.OrgStatusClosed, now, contractAddress)

	if err != nil {
		common.Logger.Error(err)
		return err
	}

	obj, err := mysqlDaoOrganizationService.getByContractAddress(contractAddress)
	if err != nil {
		return err
	}
	err = rollbackService.Insert(height, obj.Id, "t_dao_organization", obj.State, int(model.OrgStatusClosed))

	return err
}

func (mysqlDaoOrganizationService MysqlDaoOrganizationService) UpdateByContractAddress(contractAddress string, newOrgName string) error {
	now := (int)(common.NowSecond())

	sql := "update t_dao_organization set org_name = ?, update_time = ? where contract_address = ?"
	_, err := mysql.Engine.Exec(sql, newOrgName, now, contractAddress)

	if err != nil {
		common.Logger.Error(err)
		return err
	}
	return nil
}

func (mysqlDaoOrganizationService MysqlDaoOrganizationService) getByContractAddress(contractAddress string) (*models.TDaoOrganization, error) {
	obj := new(models.TDaoOrganization)
	_, err := mysql.Engine.Where("contract_address = ?", contractAddress).Get(obj)
	return obj, err
}
