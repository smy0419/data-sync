package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
)

type BtcMinerService struct{}

func (btcMinerService BtcMinerService) Insert(height int64, time int64, address string, domain string) error {
	exist, err := btcMinerService.exist(address)
	if err != nil {
		return err
	}

	if !exist {
		btcMiner := model.BtcMiner{
			Height:  height,
			Time:    time,
			Address: address,
			Domain:  domain,
		}
		_, err := mongo.MongoDB.Collection(mongo.CollectionBtcMiner).InsertOne(context.TODO(), btcMiner)
		if err != nil {
			common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionBtcMiner, err)
			return err
		}
	} else {
		err := btcMinerService.modifyDomain(height, address, domain)
		if err != nil {
			return err
		}
	}

	return nil
}

func (btcMinerService BtcMinerService) exist(address string) (bool, error) {
	exist, err := mongo.Exist(mongo.CollectionBtcMiner, bson.M{"address": address})
	return exist, err
}

func (btcMinerService BtcMinerService) getBtcMiner(address string) (*model.BtcMiner, error) {
	var result model.BtcMiner
	filter := bson.M{"address": address}
	err := mongo.FindOne(mongo.CollectionBtcMiner, filter, &result)
	return &result, err
}

func (btcMinerService BtcMinerService) modifyDomain(height int64, address string, domain string) error {
	btcMiner, err := btcMinerService.getBtcMiner(address)
	if err != nil {
		return err
	}
	if btcMiner.Domain == domain {
		return nil
	}

	filter := bson.M{"address": address}
	update := bson.M{"domain": domain}
	_, err = mongo.MongoDB.Collection(mongo.CollectionBtcMiner).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		common.Logger.Errorf("modify btc miner domain error. err: %s", err)
		return err
	}

	err = rollbackService.Insert(mongo.CollectionBtcMiner, btcMiner.ID, height, "domain", btcMiner.Domain, domain)
	if err != nil {
		common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionBtcMiner, err)
		return err
	}

	return nil
}
