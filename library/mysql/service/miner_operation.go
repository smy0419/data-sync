package service

import (
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mysql"
	models "github.com/AsimovNetwork/data-sync/library/mysql/model"
)

type MinerOperationService struct{}

var rollbackService = RollbackService{}

func (minerOperationService MinerOperationService) UpdateStatusByTxHash(txHash string, status uint8, height int64) error {
	now := (int)(common.NowSecond())

	minerOperation, err := minerOperationService.getByHash(txHash)
	if err != nil {
		return err
	}

	sql := "update t_miner_operation set tx_status = ?, update_time = ? where tx_hash = ?"
	_, err = mysql.Engine.Exec(sql, status, now, txHash)

	if err != nil {
		common.Logger.Error(err)
		return err
	}

	err = rollbackService.Insert(height, minerOperation.Id, "t_miner_operation", minerOperation.TxStatus, int(status))

	return err
}

func (minerOperationService MinerOperationService) TxHashExist(txHash string) (bool, error) {
	p := new(models.TMinerOperation)
	total, err := mysql.Engine.Where("tx_hash = ?", txHash).Count(p)
	if err != nil {
		common.Logger.Error(err)
	}
	return total > 0, err
}

func (minerOperationService MinerOperationService) getByHash(txHash string) (*models.TMinerOperation, error) {
	minerOperation := new(models.TMinerOperation)
	_, err := mysql.Engine.Where("tx_hash = ?", txHash).Get(minerOperation)
	return minerOperation, err
}
