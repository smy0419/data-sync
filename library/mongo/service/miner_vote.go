package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
)

type MinerVoteService struct{}

var minerTodoListService = MinerTodoListService{}

func (minerVoteService MinerVoteService) Insert(round int64, height int64, time int64, proposalId int64, voter string, decision bool, txHash string) error {
	vote := model.MinerVote{
		Round:      round,
		Height:     height,
		Time:       time,
		ProposalId: proposalId,
		Voter:      voter,
		Decision:   decision,
		TxHash:     txHash,
	}

	_, err := mongo.MongoDB.Collection(mongo.CollectionMinerVote).InsertOne(context.TODO(), vote)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionMinerVote, err)
		return err
	}

	// to do list done
	err = minerTodoListService.Release(height, time, round, voter, proposalId)
	if err != nil {
		common.Logger.Errorf("operate to do list from insert vote error. err: %s", err)
		return err
	}

	return nil
}
