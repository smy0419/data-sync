package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
)

type DaoVoteService struct{}

func (daoVoteService DaoVoteService) Insert(height int64, time int64, txHash string, contractAddress string, voter string, voteId int64, decision bool) error {
	insert := model.DaoVote{
		Height:          height,
		Time:            time,
		TxHash:          txHash,
		ContractAddress: contractAddress,
		Voter:           voter,
		VoteId:          voteId,
		Decision:        decision,
	}

	_, err := mongo.MongoDB.Collection(mongo.CollectionDaoVote).InsertOne(context.TODO(), insert)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionDaoVote, err)
		return err
	}
	return nil
}
