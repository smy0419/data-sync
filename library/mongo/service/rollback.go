package service

import (
	"context"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/asimov/protos"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"strings"
)

type RollbackService struct{}

var transactionService = TransactionService{}

func (rollbackService RollbackService) Rollback(height int64) error {
	// Get all collections
	collections, err := mongo.MongoDB.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	err = rollbackAddressAssetBalance(height)
	if err != nil {
		return err
	}

	err = rollbackTxCount(height)
	if err != nil {
		return err
	}

	err = rollbackTxCountCache(height)
	if err != nil {
		return err
	}

	err = rollbackTransactionHeightCache(height)
	if err != nil {
		return err
	}

	rollbackCollectionAndField, err := rollbackService.listRollbackByHeight(height)

	collectionStr := ""
	fieldStr := ""
	for _, v := range rollbackCollectionAndField {
		if strings.Index(collectionStr, v.Collection) < 0 {
			collectionStr += "\"" + v.Collection + "\","
		}
		if strings.Index(fieldStr, v.Field) < 0 {
			fieldStr += "\"" + v.Field + "\","
		}
	}
	if len(collectionStr) > 0 {
		collectionStr = collectionStr[:len(collectionStr)-1]
	}
	if len(fieldStr) > 0 {
		fieldStr = fieldStr[:len(fieldStr)-1]
	}

	rollbackList, err := rollbackService.listRollbackByCollectionAndField(height, collectionStr, fieldStr)

	for _, v := range rollbackList {
		updateFilter := bson.M{"_id": v.DocumentID}
		update := bson.M{v.Field: v.ExpectValue}

		_, err := mongo.MongoDB.Collection(v.Collection).UpdateOne(context.TODO(), updateFilter, bson.M{"$set": update})
		if err != nil {
			common.Logger.Errorf("rollback collection %s error. err: %s", v.Collection, err)
			return err
		}
	}

	deleteFilter := bson.M{
		"height": bson.M{"$gt": height},
	}
	for _, v := range collections {
		_, err = mongo.MongoDB.Collection(v).DeleteMany(context.TODO(), deleteFilter)
		if err != nil {
			common.Logger.Errorf("delete from %s error. err: %s", v, err)
			return err
		}
	}

	_, err = mongo.MongoDB.Collection(mongo.CollectionRollback).DeleteMany(context.TODO(), deleteFilter)
	if err != nil {
		common.Logger.Errorf("delete from %s error. err: %s", mongo.CollectionRollback, err)
		return err
	}

	return nil
}

func (rollbackService RollbackService) Insert(
	collection string,
	documentId primitive.ObjectID,
	height int64,
	field string,
	originalValue interface{},
	expectValue interface{}) error {

	rollback := model.Rollback{
		Collection:    collection,
		DocumentID:    documentId,
		Height:        height,
		Time:          common.Now(),
		Field:         field,
		OriginalValue: originalValue,
		ExpectValue:   expectValue,
	}

	_, err := mongo.MongoDB.Collection(mongo.CollectionRollback).InsertOne(context.TODO(), rollback)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionRollback, err)
		return err
	}
	return nil
}

func (rollbackService RollbackService) ListRollback(height int64) ([]*model.Rollback, error) {
	filter := bson.M{
		"height": bson.M{"$gte": height},
	}

	results, err := mongo.Find(mongo.CollectionRollback, filter, reflect.TypeOf(model.Rollback{}), reflect.TypeOf(&model.Rollback{}))
	return results.([]*model.Rollback), err
}

func (rollbackService RollbackService) listRollbackByHeight(height int64) ([]*model.Rollback, error) {
	var commandByte []byte
	command := `[
					{
						"$match": {
							"height": {
								"$gt": %d
							}
						}
					},
					{
						"$group": {
							"_id": {
								"field": "$field",
								"collection": "$collection"
							}
						}
					}
    			]`
	commandByte = []byte(fmt.Sprintf(command, height))

	pipeLine := mongoDriver.Pipeline{}
	err := bson.UnmarshalExtJSON([]byte(commandByte), true, &pipeLine)
	if err != nil {
		return nil, err
	}

	cursor, err := mongo.MongoDB.Collection(mongo.CollectionRollback).Aggregate(context.TODO(), pipeLine)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Rollback, 0)
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var m bson.M
		err = cursor.Decode(&m)
		if err != nil {
			return nil, err
		}
		tmp := m["_id"].(primitive.M)
		rollback := &model.Rollback{
			Collection: tmp["collection"].(string),
			Field:      tmp["field"].(string),
		}
		result = append(result, rollback)
	}
	return result, nil
}

func (rollbackService RollbackService) listRollbackByCollectionAndField(height int64, collectionStr string, fieldStr string) ([]*model.Rollback, error) {
	var commandByte []byte
	command := `[
			 		{
					    "$match": {
							"height": {
								"$lte": %d
							},
							"collection": {
								"$in": [%s]
							},
							"field": {
								"$in": [%s]
							}
						}
				 	},
					{
				 		"$sort": {
							"time": -1
						}
				 	},
				 	{
						"$group": {
							"_id": {
				 				"field": "$field",
								"collection": "$collection"
							},
							"data": {
								"$first": "$$ROOT"
							}
						}
					}
				]`
	commandByte = []byte(fmt.Sprintf(command, height, collectionStr, fieldStr))

	pipeLine := mongoDriver.Pipeline{}
	err := bson.UnmarshalExtJSON([]byte(commandByte), true, &pipeLine)
	if err != nil {
		return nil, err
	}

	cursor, err := mongo.MongoDB.Collection(mongo.CollectionRollback).Aggregate(context.TODO(), pipeLine)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Rollback, 0)
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var m bson.M
		err = cursor.Decode(&m)
		if err != nil {
			return nil, err
		}
		tmp := m["data"].(primitive.M)
		rollback := &model.Rollback{
			DocumentID:    tmp["document_id"].(primitive.ObjectID),
			Height:        tmp["height"].(int64),
			Time:          tmp["time"].(int64),
			OriginalValue: tmp["original_value"],
			ExpectValue:   tmp["expect_value"],
			Collection:    tmp["collection"].(string),
			Field:         tmp["field"].(string),
		}
		result = append(result, rollback)
	}
	return result, nil
}

func rollbackTxCount(height int64) error {
	if height <= 0 {
		_, err := mongo.MongoDB.Collection(mongo.CollectionTransactionCount).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return err
		}
	}

	filter := bson.M{
		"height": bson.M{"$gt": height},
	}

	update := bson.M{
		"$inc": bson.M{"tx_count": -1},
	}

	transactions, err := mongo.Find(mongo.CollectionTransaction, filter, reflect.TypeOf(model.Transaction{}), reflect.TypeOf(&model.Transaction{}))
	if err != nil {
		return err
	}

	for _, transaction := range transactions.([]*model.Transaction) {
		vins := transaction.Vin
		vouts := transaction.Vout
		keySlice := make([]string, 0)
		for _, vin := range vins {
			if vin.PrevOut != nil {
				// append asset
				keySlice = append(keySlice, vin.PrevOut.Asset)
				for _, address := range vin.PrevOut.Addresses {
					if address[:4] == common.CitizenPrefix {
						// append normal address
						keySlice = append(keySlice, address)
					}
				}
			}
		}

		for _, vout := range vouts {
			for _, address := range vout.ScriptPubKey.Addresses {
				// append normal address or contract address
				if address[:4] == common.CitizenPrefix || address[:4] == common.ContractPrefix {
					keySlice = append(keySlice, address)
				}
			}
		}

		keySlice = common.RemoveRepeatByLoop(keySlice)

		if len(keySlice) > 0 {
			for _, key := range keySlice {
				updateFilter := bson.M{
					"key": key,
				}

				_, err := mongo.MongoDB.Collection(mongo.CollectionTransactionCount).UpdateOne(context.TODO(), updateFilter, update)
				if err != nil {
					common.Logger.Errorf("rollback %s error. err: %s", mongo.CollectionTransactionCount, err)
					return err
				}
			}
		}
	}

	return nil
}

func rollbackAddressAssetBalance(height int64) error {
	if height <= 0 {
		_, err := mongo.MongoDB.Collection(mongo.CollectionAddressAssetBalance).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return err
		}
	}

	filter := bson.M{
		"height": bson.M{"$gt": height},
	}

	transactions, err := mongo.Find(mongo.CollectionTransaction, filter, reflect.TypeOf(model.Transaction{}), reflect.TypeOf(&model.Transaction{}))
	if err != nil {
		return err
	}
	for _, v := range transactions.([]*model.Transaction) {
		// increase address_asset_balance
		for _, vin := range v.Vin {
			if vin.PrevOut != nil {
				for _, address := range vin.PrevOut.Addresses {
					assetObj := protos.AssetFromBytes(asimovCommon.Hex2Bytes(vin.PrevOut.Asset))
					updateFilter := bson.M{"address": address, "asset": vin.PrevOut.Asset}
					update := bson.M{}
					if assetObj.IsIndivisible() {
						update = bson.M{
							"$inc": bson.M{"balance": 1},
						}
					} else {
						update = bson.M{
							"$inc": bson.M{"balance": vin.PrevOut.Value},
						}
					}

					_, err = mongo.MongoDB.Collection(mongo.CollectionAddressAssetBalance).UpdateOne(context.TODO(), updateFilter, update)
					if err != nil {
						common.Logger.Errorf("rollback address_asset_balance error. err: %s", err)
						return err
					}
				}
			}
		}

		// decrease address_asset_balance
		for _, vout := range v.Vout {
			for _, address := range vout.ScriptPubKey.Addresses {
				assetObj := protos.AssetFromBytes(asimovCommon.Hex2Bytes(vout.Asset))
				updateFilter := bson.M{"address": address, "asset": vout.Asset}
				update := bson.M{}
				if assetObj.IsIndivisible() {
					update = bson.M{
						"$inc": bson.M{"balance": -1},
					}
				} else {
					update = bson.M{
						"$inc": bson.M{"balance": -vout.Value},
					}
				}

				_, err = mongo.MongoDB.Collection(mongo.CollectionAddressAssetBalance).UpdateOne(context.TODO(), updateFilter, update)
				if err != nil {
					common.Logger.Errorf("rollback address_asset_balance error. err: %s", err)
					return err
				}
			}
		}
	}

	return nil
}

func rollbackTxCountCache(height int64) error {
	rollbackTxCount, err := blockService.CountTransaction(height)
	if err != nil {
		return err
	}
	return transactionCacheService.RollbackTxCount(0 - rollbackTxCount)
}

func rollbackTransactionHeightCache(height int64) error {
	transactions, err := transactionService.GetTransactionsGreaterThanHeight(height)
	if err != nil {
		return err
	}
	txHash := make([]interface{}, 0)
	for _, v := range transactions {
		txHash = append(txHash, v.Hash)
	}

	return transactionCacheService.RollbackTransactionHeight(txHash)
}
