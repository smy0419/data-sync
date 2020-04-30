package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
)

type FoundationBalanceSheetService struct {
}

func (foundationBalanceSheetService FoundationBalanceSheetService) Insert(height int64, time int64, txHash string, address string, amount int64, asset string, transferType uint8, proposalType int8) error {
	insert := model.FoundationBalanceSheet{
		Address:      address,
		Height:       height,
		Time:         time,
		Asset:        asset,
		Amount:       amount,
		TransferType: transferType,
		ProposalType: proposalType,
		TxHash:       txHash,
	}

	_, err := mongo.MongoDB.Collection(mongo.CollectionFoundationBalanceSheet).InsertOne(context.TODO(), insert)
	if err != nil {
		common.Logger.Errorf("insert foundation balance sheet error. err: %s", err)
		return err
	}
	return nil
}
