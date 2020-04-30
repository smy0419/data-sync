package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
)

type ValidatorRelationService struct{}

func (validatorRelationService ValidatorRelationService) Insert(height int64, time int64, btcMinerAddress string, address string) error {
	exist, err := validatorRelationService.exist(address)
	if err != nil {
		return err
	}
	if exist {
		err = validatorRelationService.release(height, time, btcMinerAddress)
		if err != nil {
			return err
		}
	}

	validatorRelation := model.ValidatorRelation{
		Height:          height,
		Time:            time,
		BtcMinerAddress: btcMinerAddress,
		Bind:            true,
		Address:         address,
	}
	_, err = mongo.MongoDB.Collection(mongo.CollectionValidatorRelation).InsertOne(context.TODO(), validatorRelation)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionValidatorRelation, err)
		return err
	}

	return nil
}

func (validatorRelationService ValidatorRelationService) release(height int64, time int64, btcMinerAddress string) error {
	validatorRelation, err := validatorRelationService.getValidatorRelation(btcMinerAddress)
	if err != nil {
		return err
	}

	filter := bson.M{"btc_miner_address": btcMinerAddress, "bind": true}
	update := bson.M{"bind": false}
	_, err = mongo.MongoDB.Collection(mongo.CollectionValidatorRelation).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		common.Logger.Errorf("release validator relation error. err: %s", err)
		return err
	}

	err = rollbackService.Insert(mongo.CollectionValidatorRelation, validatorRelation.ID, height, "bind", validatorRelation.Bind, false)
	if err != nil {
		common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionValidatorRelation, err)
		return err
	}

	return nil
}

func (validatorRelationService ValidatorRelationService) exist(btcMinerAddress string) (bool, error) {
	exist, err := mongo.Exist(mongo.CollectionValidatorRelation, bson.M{"btc_miner_address": btcMinerAddress, "bind": true})
	return exist, err
}

func (validatorRelationService ValidatorRelationService) getValidatorRelation(btcMinerAddress string) (*model.ValidatorRelation, error) {
	var result model.ValidatorRelation
	filter := bson.M{"btc_miner_address": btcMinerAddress, "bind": true}
	err := mongo.FindOne(mongo.CollectionValidatorRelation, filter, &result)
	return &result, err
}
