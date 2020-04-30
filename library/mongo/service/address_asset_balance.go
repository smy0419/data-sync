package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

type AddressAssetBalanceService struct{}

//func (addressAssetBalanceService AddressAssetBalanceService) InsertOrUpdate(height int64, addressAssetBalanceSlice []model.AddressAssetBalance) error {
//	//var recordInMongo model.AddressAssetBalance
//	for _, v := range addressAssetBalanceSlice {
//		if v.Balance == 0 {
//			continue
//		}
//		filter := bson.M{
//			"address": v.Address,
//			"asset":   v.Asset,
//		}
//
//		update := bson.M{
//			"$inc": bson.M{"balance": v.Balance},
//		}
//		var flag = true
//		updateOptions := options.UpdateOptions{
//			Upsert: &flag,
//		}
//		_, err := mongo.MongoDB.Collection(mongo.CollectionAddressAssetBalance).UpdateOne(context.TODO(), filter, update, &updateOptions)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (addressAssetBalanceService AddressAssetBalanceService) InsertOrUpdate(height int64, addressAssetBalanceSlice []model.AddressAssetBalance) error {
	var operations []mongoDriver.WriteModel
	//var recordInMongo model.AddressAssetBalance
	for _, v := range addressAssetBalanceSlice {
		if v.Balance == 0 {
			continue
		}
		filter := bson.M{
			"address": v.Address,
			"asset":   v.Asset,
		}
		update := bson.M{
			"$inc": bson.M{"balance": v.Balance},
		}
		operation := mongoDriver.NewUpdateOneModel()
		operation.SetFilter(filter)
		operation.SetUpdate(update)
		operation.SetUpsert(true)
		operations = append(operations, operation)
	}

	if len(operations) > 0 {
		_, err := mongo.MongoDB.Collection(mongo.CollectionAddressAssetBalance).BulkWrite(
			context.Background(),
			operations,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
