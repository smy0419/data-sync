package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"github.com/AsimovNetwork/data-sync/library/response"
	"go.mongodb.org/mongo-driver/bson"
)

type MinerSignUpService struct{}

func (minerSignUpService MinerSignUpService) Insert(height int64, time int64, txHash string, round int64, address string) error {
	validator, err := validatorService.GetValidator(address)
	if err != nil {
		return err
	}

	insert := model.MinerSignUp{
		Height:     height,
		Time:       time,
		TxHash:     txHash,
		Round:      round + 1,
		Address:    address,
		Produced:   validator.ActualBlocks,
		Planed:     validator.PlannedBlocks,
		Efficiency: common.CalculateEfficiency(validator.ActualBlocks, validator.PlannedBlocks),
	}

	_, err = mongo.MongoDB.Collection(mongo.CollectionMinerSignUp).InsertOne(context.TODO(), insert)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionMinerSignUp, err)
		return err
	}
	return nil
}

func (minerSignUpService MinerSignUpService) UpdateBlocksAndEfficiency(height int64, round int64, validatorSlice []string, plannedBlocks []int64, actualBlocks []int64) error {
	if len(validatorSlice) != len(plannedBlocks) || len(plannedBlocks) != len(actualBlocks) {
		errMsg := fmt.Sprintf("invalid parameter, length of three parameters should be equivalent")
		common.Logger.Error(errMsg)
		return errors.New(errMsg)
	}

	for i, v := range validatorSlice {
		filter := bson.M{
			"round":   round + 1,
			"address": v,
		}
		var signUp model.MinerSignUp
		err := mongo.FindOne(mongo.CollectionMinerSignUp, filter, &signUp)
		if err != nil {
			if response.IsDataNotExistError(err) {
				return nil
			}
			return err
		}

		update := bson.M{
			"produced":   signUp.Produced + actualBlocks[i],
			"planed":     signUp.Planed + plannedBlocks[i],
			"efficiency": common.CalculateEfficiency(signUp.Produced+actualBlocks[i], signUp.Planed+plannedBlocks[i]),
		}

		_, err = mongo.MongoDB.Collection(mongo.CollectionMinerSignUp).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
		if err != nil {
			common.Logger.Errorf("change miner sign up efficiency or produced error. err: %s", err)
			return err
		}

		err = rollbackService.Insert(mongo.CollectionMinerSignUp, signUp.ID, height, "efficiency", signUp.Efficiency, common.CalculateEfficiency(signUp.Produced+actualBlocks[i], signUp.Planed+plannedBlocks[i]))
		if err != nil {
			common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionMinerSignUp, err)
			return err
		}

		err = rollbackService.Insert(mongo.CollectionMinerSignUp, signUp.ID, height, "produced", signUp.Produced, actualBlocks[i])
		if err != nil {
			common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionMinerSignUp, err)
			return err
		}

		err = rollbackService.Insert(mongo.CollectionMinerSignUp, signUp.ID, height, "planed", signUp.Planed, plannedBlocks[i])
		if err != nil {
			common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionMinerSignUp, err)
			return err
		}
	}
	return nil
}
