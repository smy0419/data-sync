package service

import (
	"context"
	"errors"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
)

type ValidatorService struct{}

var emptyLocation = model.Location{}

func (validatorService ValidatorService) Insert(height int64, time int64, address string) error {
	exist, err := validatorService.Exist(address)
	if err != nil {
		return err
	}

	if !exist {
		insert := model.Validator{
			Height:        height,
			Time:          time,
			Address:       address,
			Location:      emptyLocation,
			PlannedBlocks: 0,
			ActualBlocks:  0,
		}
		_, err := mongo.MongoDB.Collection(mongo.CollectionValidator).InsertOne(context.TODO(), insert)
		if err != nil {
			common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionValidator, err)
			return err
		}
	}

	return nil
}

func (validatorService ValidatorService) Exist(address string) (bool, error) {
	exist, err := mongo.Exist(mongo.CollectionValidator, bson.M{"address": address})
	return exist, err
}

func (validatorService ValidatorService) GetValidator(address string) (*model.Validator, error) {
	var result model.Validator
	filter := bson.M{"address": address}
	err := mongo.FindOne(mongo.CollectionValidator, filter, &result)
	return &result, err
}

func (validatorService ValidatorService) ModifyBlocks(height int64, addresses []string, plannedBlocks []int64, actualBlocks []int64) error {
	if len(addresses) != len(plannedBlocks) || len(plannedBlocks) != len(actualBlocks) {
		return errors.New("invalid parameter, length of three parameters should be equivalent")
	}

	for i, v := range addresses {
		filter := bson.M{
			"address": v,
		}

		validator, err := validatorService.GetValidator(v)
		if err != nil {
			return err
		}

		update := bson.M{
			"planned_blocks": validator.PlannedBlocks + plannedBlocks[i],
			"actual_blocks":  validator.ActualBlocks + actualBlocks[i],
		}

		_, err = mongo.MongoDB.Collection(mongo.CollectionValidator).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
		if err != nil {
			common.Logger.Errorf("change validator blocks error. err: %s", err)
			return err
		}

		err = rollbackService.Insert(mongo.CollectionValidator, validator.ID, height, "planned_blocks", validator.PlannedBlocks, validator.PlannedBlocks+plannedBlocks[i])
		if err != nil {
			common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionValidator, err)
			return err
		}
		err = rollbackService.Insert(mongo.CollectionValidator, validator.ID, height, "actual_blocks", validator.ActualBlocks, validator.ActualBlocks+actualBlocks[i])
		if err != nil {
			common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionValidator, err)
			return err
		}
	}
	return nil
}
