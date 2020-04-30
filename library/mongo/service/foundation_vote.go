package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
)

type FoundationVoteService struct{}

func (foundationVoteService FoundationVoteService) Insert(height int64, time int64, proposalId int64, voter string, decision bool, txHash string) error {
	insert := model.FoundationVote{
		Height:     height,
		Time:       time,
		ProposalId: proposalId,
		Voter:      voter,
		Decision:   decision,
		TxHash:     txHash,
	}

	_, err := mongo.MongoDB.Collection(mongo.CollectionFoundationVote).InsertOne(context.TODO(), insert)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionFoundationVote, err)
		return err
	}

	// to do list done
	err = foundationTodoListService.Release(height, voter, proposalId)
	if err != nil {
		common.Logger.Errorf("operate to do list from insert vote error. err: %s", err)
		return err
	}

	return nil
}
