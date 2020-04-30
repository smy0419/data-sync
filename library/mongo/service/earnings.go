package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EarningService struct{}

func (earningService EarningService) Insert(height int64, time int64, txHash string, earnings map[string]map[string]int64) error {
	for k, v := range earnings {
		earning := model.Earning{
			Time:    time,
			Height:  height,
			TxHash:  txHash,
			Address: k,
		}

		insertResult, err := mongo.MongoDB.Collection(mongo.CollectionEarning).InsertOne(context.TODO(), earning)
		if err != nil {
			common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionEarning, err)
			return err
		}

		id := insertResult.InsertedID.(primitive.ObjectID)
		for asset, amount := range v {
			err := earningService.insertEarningAsset(height, time, id, asset, amount)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (earningService EarningService) insertEarningAsset(height int64, time int64, id primitive.ObjectID, asset string, amount int64) error {
	earningAsset := model.EarningAsset{
		Time:      time,
		Height:    height,
		EarningId: id,
		Asset:     asset,
		Value:     amount,
	}

	_, err := mongo.MongoDB.Collection(mongo.CollectionEarningAsset).InsertOne(context.TODO(), earningAsset)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionEarningAsset, err)
		return err
	}

	return nil
}
