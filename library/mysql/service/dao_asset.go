package service

import (
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"github.com/AsimovNetwork/data-sync/library/mysql"
)

type MysqlDaoAssetService struct{}

func (mysqlDaoAssetService MysqlDaoAssetService) UpdateByTxHash(txHash string, asset string) error {
	now := (int)(common.NowSecond())

	sql := "update t_dao_asset set asset = ?, update_time = ? where tx_hash = ?"
	_, err := mysql.Engine.Exec(sql, asset, now, txHash)

	if err != nil {
		common.Logger.Error(err)
		return err
	}

	return nil
}

func (mysqlDaoAssetService MysqlDaoAssetService) UpdateByAsset(asset string) error {
	now := (int)(common.NowSecond())

	sql := "update t_dao_asset set asset_status = ?, update_time = ? where asset = ?"
	_, err := mysql.Engine.Exec(sql, model.AssetStatusSuccess, now, asset)

	if err != nil {
		common.Logger.Error(err)
		return err
	}

	return nil
}
