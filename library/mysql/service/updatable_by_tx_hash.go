package service

import (
	"github.com/AsimovNetwork/asimov/rpcs/rpcjson"
	"github.com/AsimovNetwork/data-sync/library/mysql/constant"
)

type UpdatableByTxhash interface {
	TxHashExist(txHash string) (bool, error)
	UpdateStatusByTxHash(txHash string, status uint8, height int64) error
}

var updatableHandler = []UpdatableByTxhash{MinerOperationService{}, FoundationOperationService{}, DaoOperationService{}}

type UpdatableService struct{}

func (updatableService UpdatableService) HandleUpdateStatus(height int64, receipts []*rpcjson.ReceiptResult) error {
	for _, receipt := range receipts {
		for _, handler := range updatableHandler {
			exist, err := handler.TxHashExist(receipt.TxHash)
			if err != nil {
				return err
			}
			if !exist {
				continue
			}
			var status uint8
			if receipt.Status > 0 {
				status = constant.SuccessStatus
			} else {
				status = constant.FailedStatus
			}
			err = handler.UpdateStatusByTxHash(receipt.TxHash, status, height)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
