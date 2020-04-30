package service

import (
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"github.com/AsimovNetwork/data-sync/library/mysql"
)

type MysqlDaoIndivisibleAssetService struct{}

func (mysqlDaoIndivisibleAssetService MysqlDaoIndivisibleAssetService) UpdateByAsset(asset string) error {
	now := (int)(common.NowSecond())

	sql := "update t_dao_indivisible_asset set asset_status = ?, update_time = ? where asset = ?"
	_, err := mysql.Engine.Exec(sql, model.AssetStatusSuccess, now, asset)

	if err != nil {
		common.Logger.Error(err)
		return err
	}

	return nil
}

func (mysqlDaoIndivisibleAssetService MysqlDaoIndivisibleAssetService) UpdateByAssetAndVoucherId(asset string, voucherId int64) error {
	now := (int)(common.NowSecond())

	sql := "update t_dao_indivisible_asset set asset_status = ?, update_time = ? where asset = ? and voucher_id = ?"
	_, err := mysql.Engine.Exec(sql, model.AssetStatusSuccess, now, asset, voucherId)

	if err != nil {
		common.Logger.Error(err)
		return err
	}

	return nil
}
