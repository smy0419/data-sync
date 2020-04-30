package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"gopkg.in/mgo.v2/bson"
)

type DaoOrganizationAssetService struct{}

func (daoOrganizationAssetService DaoOrganizationAssetService) Insert(height int64, time int64, contractAddress string, asset string, amount int64, assetType uint32, assetIndex uint32) error {
	insert := model.DaoOrganizationAsset{
		Height:          height,
		Time:            time,
		ContractAddress: contractAddress,
		Asset:           asset,
		AssetType:       assetType,
		AssetIndex:      assetIndex,
	}
	_, err := mongo.MongoDB.Collection(mongo.CollectionDaoOrganizationAsset).InsertOne(context.TODO(), insert)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionDaoOrganizationAsset, err)
		return err
	}
	return nil
}

func (daoOrganizationAssetService DaoOrganizationAssetService) GetAsset(contractAddress string, assetIndex uint32) (*model.DaoOrganizationAsset, error) {
	var result model.DaoOrganizationAsset
	filter := bson.M{
		"contract_address": contractAddress,
		"asset_index":      assetIndex,
	}
	err := mongo.FindOne(mongo.CollectionDaoOrganizationAsset, filter, &result)
	return &result, err
}
