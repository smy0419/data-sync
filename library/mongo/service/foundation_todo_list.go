package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
)

type FoundationTodoListService struct{}

var foundationProposalService = FoundationProposalService{}

func (foundationTodoListService FoundationTodoListService) Insert(height int64, time int64, operators []string, todoId int64) error {
	var proposal *model.FoundationProposal
	proposal, err := foundationProposalService.GetProposal(todoId)
	if err != nil {
		return err
	}

	inserts := make([]interface{}, 0)
	for _, v := range operators {
		todo := model.FoundationTodoList{
			Height:       height,
			Time:         time,
			Operator:     v,
			TodoId:       todoId,
			ProposalType: proposal.ProposalType,
			Operated:     false,
			EndTime:      proposal.EndTime,
		}
		inserts = append(inserts, todo)
	}

	_, err = mongo.MongoDB.Collection(mongo.CollectionFoundationTodoList).InsertMany(context.TODO(), inserts)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionFoundationTodoList, err)
		return err
	}

	return nil
}

func (foundationTodoListService FoundationTodoListService) Release(height int64, operator string, todoId int64) error {
	filter := bson.M{
		"todo_id":  todoId,
		"operator": operator,
	}

	update := bson.M{
		"operated": true,
	}

	var todo model.FoundationTodoList
	err := mongo.FindOne(mongo.CollectionFoundationTodoList, filter, &todo)
	if err != nil {
		return err
	}

	err = rollbackService.Insert(mongo.CollectionFoundationTodoList, todo.ID, height, "operated", todo.Operated, true)
	if err != nil {
		common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionFoundationTodoList, err)
		return err
	}

	_, err = mongo.MongoDB.Collection(mongo.CollectionFoundationTodoList).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	return err
}
