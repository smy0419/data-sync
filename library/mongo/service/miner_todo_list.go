package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
)

type MinerTodoListService struct{}

var minerProposalService = MinerProposalService{}

func (minerTodoListService MinerTodoListService) Insert(height int64, time int64, round int64, operators []string, todoId int64) error {
	var minerProposal *model.MinerProposal
	var err error
	minerProposal, err = minerProposalService.GetProposal(todoId)
	if err != nil {
		return err
	}

	inserts := make([]interface{}, 0)
	for _, v := range operators {
		insert := model.MinerTodoList{
			Height:     height,
			Time:       time,
			Round:      round,
			Operator:   v,
			ActionId:   todoId,
			ActionType: minerProposal.Type,
			EndTime:    minerProposal.EndTime,
			Operated:   false,
		}
		inserts = append(inserts, insert)
	}

	_, err = mongo.MongoDB.Collection(mongo.CollectionMinerTodoList).InsertMany(context.TODO(), inserts)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionMinerTodoList, err)
		return err
	}

	return nil
}

func (minerTodoListService MinerTodoListService) Release(height int64, time int64, round int64, operator string, todoId int64) error {
	filter := bson.M{
		"round":     round,
		"operator":  operator,
		"action_id": todoId,
	}

	var todoList model.MinerTodoList
	err := mongo.FindOne(mongo.CollectionMinerTodoList, filter, &todoList)
	if err != nil {
		return err
	}

	update := bson.M{
		"operated": true,
	}
	_, err = mongo.MongoDB.Collection(mongo.CollectionMinerTodoList).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return err
	}
	err = rollbackService.Insert(mongo.CollectionMinerTodoList, todoList.ID, height, "operated", todoList.Operated, true)
	if err != nil {
		common.Logger.Errorf("insert rollback  error. collection: %s, err: %s", mongo.CollectionMinerTodoList, err)
		return err
	}

	return err
}
