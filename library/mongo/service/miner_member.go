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

type MinerMemberService struct{}

func (minerMemberService MinerMemberService) Insert(height int64, time int64, round int64, addresses []string) error {
	inserts := make([]interface{}, 0)
	for _, v := range addresses {
		var produced int64 = 0
		var planed int64 = 0
		var efficiency int32 = 0

		validator, err := validatorService.GetValidator(v)
		if err != nil && !response.IsDataNotExistError(err) {
			return err
		}
		if err == nil {
			produced = validator.ActualBlocks
			planed = validator.PlannedBlocks
			efficiency = common.CalculateEfficiency(validator.ActualBlocks, validator.PlannedBlocks)
		}

		member := model.MinerMember{
			Height:     height,
			Time:       time,
			Round:      round,
			Address:    v,
			Produced:   produced,
			Planed:     planed,
			Efficiency: efficiency,
		}

		inserts = append(inserts, member)
	}
	_, err := mongo.MongoDB.Collection(mongo.CollectionMinerMember).InsertMany(context.TODO(), inserts)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionMinerMember, err)
		return err
	}
	return nil
}

func (minerMemberService MinerMemberService) UpdateBlocksAndEfficiency(height int64, round int64, validatorSlice []string, plannedBlocks []int64, actualBlocks []int64) error {
	if len(validatorSlice) != len(plannedBlocks) || len(plannedBlocks) != len(actualBlocks) {
		errMsg := fmt.Sprintf("invalid parameter, length of three parameters should be equivalent")
		common.Logger.Error(errMsg)
		return errors.New(errMsg)
	}

	for i, v := range validatorSlice {
		filter := bson.M{
			"round":   round,
			"address": v,
		}

		var member model.MinerMember
		err := mongo.FindOne(mongo.CollectionMinerMember, filter, &member)
		if err != nil {
			if response.IsDataNotExistError(err) {
				return nil
			}
			return err
		}

		update := bson.M{
			"produced":   member.Produced + actualBlocks[i],
			"planed":     member.Planed + plannedBlocks[i],
			"efficiency": common.CalculateEfficiency(member.Produced+actualBlocks[i], member.Planed+plannedBlocks[i]),
		}

		_, err = mongo.MongoDB.Collection(mongo.CollectionMinerMember).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
		if err != nil {
			common.Logger.Errorf("change miner member efficiency or produced error. err: %s", err)
			return err
		}

		err = rollbackService.Insert(mongo.CollectionMinerMember, member.ID, height, "efficiency", member.Efficiency, common.CalculateEfficiency(member.Produced+actualBlocks[i], member.Planed+plannedBlocks[i]))
		if err != nil {
			common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionMinerMember, err)
			return err
		}

		err = rollbackService.Insert(mongo.CollectionMinerMember, member.ID, height, "produced", member.Produced, member.Produced+actualBlocks[i])
		if err != nil {
			common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionMinerMember, err)
			return err
		}

		err = rollbackService.Insert(mongo.CollectionMinerMember, member.ID, height, "planed", member.Planed, member.Planed+plannedBlocks[i])
		if err != nil {
			common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionMinerMember, err)
			return err
		}
	}
	return nil
}
