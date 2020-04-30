package service

import (
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mysql"
	models "github.com/AsimovNetwork/data-sync/library/mysql/model"
)

type FoundationOperationService struct{}

func (foundationOperationService FoundationOperationService) UpdateStatusByTxHash(txHash string, status uint8, height int64) error {
	now := (int)(common.NowSecond())

	foundationOperation, err := foundationOperationService.getByHash(txHash)
	if err != nil {
		return err
	}

	sql := "update t_foundation_operation set tx_status = ?, update_time = ? where tx_hash = ?"
	_, err = mysql.Engine.Exec(sql, status, now, txHash)

	if err != nil {
		common.Logger.Error(err)
		return err
	}

	err = rollbackService.Insert(height, foundationOperation.Id, "t_foundation_operation", foundationOperation.TxStatus, int(status))

	return err
}

func (foundationOperationService FoundationOperationService) TxHashExist(txHash string) (bool, error) {
	p := new(models.TFoundationOperation)
	total, err := mysql.Engine.Where("tx_hash = ?", txHash).Count(p)
	if err != nil {
		common.Logger.Error(err)
	}
	return total > 0, err
}

func (foundationOperationService FoundationOperationService) getByHash(txHash string) (*models.TFoundationOperation, error) {
	foundationOperation := new(models.TFoundationOperation)
	_, err := mysql.Engine.Where("tx_hash = ?", txHash).Get(foundationOperation)
	return foundationOperation, err
}
