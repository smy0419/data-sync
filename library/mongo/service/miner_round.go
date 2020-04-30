package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
)

type MinerRoundService struct{}

func (minerRoundService MinerRoundService) Insert(
	height int64,
	time int64,
	round int64,
	startTime int64,
	endTime int64,
) error {
	r := model.MinerRound{
		Height:    height,
		Time:      time,
		Round:     round,
		StartTime: startTime,
		EndTime:   endTime,
	}
	_, err := mongo.MongoDB.Collection(mongo.CollectionMinerRound).InsertOne(context.TODO(), r)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionMinerRound, err)
		return err
	}
	return nil
}
