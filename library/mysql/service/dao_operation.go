package service

import (
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mysql"
	"github.com/AsimovNetwork/data-sync/library/mysql/constant"
	models "github.com/AsimovNetwork/data-sync/library/mysql/model"
)

type DaoOperationService struct{}

func (daoOperationService DaoOperationService) UpdateStatusByTxHash(txHash string, status uint8, height int64) error {
	now := (int)(common.NowSecond())

	operation, err := daoOperationService.getByHash(txHash)
	if err != nil {
		return err
	}

	sql := "update t_dao_operation set tx_status = ?, update_time = ? where tx_hash = ?"
	_, err = mysql.Engine.Exec(sql, status, now, txHash)

	if err != nil {
		common.Logger.Error(err)
		return err
	}

	// if transaction failed and operation is issue asset or mint asset
	// update t_dao_asset and t_dao_indivisible_asset set asset_status to failure
	if status == constant.FailedStatus && (operation.OperationType == constant.OperationTypeIssueAsset || operation.OperationType == constant.OperationTypeMintAsset) {
		assetSql := "update t_dao_asset set asset_status = ?, update_time = ? where tx_hash = ?"
		asset, err := daoOperationService.getAssetByHash(txHash)
		if err != nil {
			return err
		}
		_, err = mysql.Engine.Exec(assetSql, constant.FailedStatus, now, txHash)
		if err != nil {
			common.Logger.Error(err)
			return err
		}

		indivisibleSql := "update t_dao_indivisible_asset set asset_status = ?, update_time = ? where tx_hash = ?"
		indivisibleAsset, err := daoOperationService.getIndivisibleAssetByHash(txHash)
		if err != nil {
			return err
		}
		_, err = mysql.Engine.Exec(indivisibleSql, constant.FailedStatus, now, txHash)
		if err != nil {
			common.Logger.Error(err)
			return err
		}

		err = rollbackService.Insert(height, asset.Id, "t_dao_asset", constant.InitStatus, constant.FailedStatus)
		if err != nil {
			common.Logger.Error(err)
			return err
		}

		err = rollbackService.Insert(height, indivisibleAsset.Id, "t_dao_indivisible_asset", constant.InitStatus, constant.FailedStatus)
		if err != nil {
			common.Logger.Error(err)
			return err
		}
	}

	err = rollbackService.Insert(height, operation.Id, "t_dao_operation", operation.TxStatus, int(status))
	return err
}

func (daoOperationService DaoOperationService) TxHashExist(txHash string) (bool, error) {
	p := new(models.TDaoOperation)
	total, err := mysql.Engine.Where("tx_hash = ?", txHash).Count(p)
	if err != nil {
		common.Logger.Error(err)
	}
	return total > 0, err
}

func (daoOperationService DaoOperationService) getByHash(txHash string) (*models.TDaoOperation, error) {
	operation := new(models.TDaoOperation)
	_, err := mysql.Engine.Where("tx_hash = ?", txHash).Get(operation)
	return operation, err
}

func (daoOperationService DaoOperationService) getAssetByHash(txHash string) (*models.TDaoAsset, error) {
	asset := new(models.TDaoAsset)
	_, err := mysql.Engine.Where("tx_hash = ?", txHash).Get(asset)
	return asset, err
}

func (daoOperationService DaoOperationService) getIndivisibleAssetByHash(txHash string) (*models.TDaoIndivisibleAsset, error) {
	indivisibleAsset := new(models.TDaoIndivisibleAsset)
	_, err := mysql.Engine.Where("tx_hash = ?", txHash).Get(indivisibleAsset)
	return indivisibleAsset, err
}
