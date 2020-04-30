package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
)

type FoundationProposalService struct{}

var foundationTodoListService = FoundationTodoListService{}

func (foundationProposalService FoundationProposalService) Insert(height int64, time int64, proposalId int64, proposalType uint8, proposer string, endTime int64, txHash string) error {
	insert := model.FoundationProposal{
		Height:       height,
		Time:         time,
		EndTime:      endTime,
		ProposalId:   proposalId,
		Address:      proposer,
		ProposalType: proposalType,
		Status:       model.ProposalStatusOnGoing,
		TxHash:       txHash,
	}

	_, err := mongo.MongoDB.Collection(mongo.CollectionFoundationProposal).InsertOne(context.TODO(), insert)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionFoundationProposal, err)
		return err
	}

	return nil
}

func (foundationProposalService FoundationProposalService) ChangeStatus(height int64, time int64, proposalId int64, status uint8) error {
	filter := bson.M{"proposal_id": proposalId}
	update := bson.M{"status": status}
	_, err := mongo.MongoDB.Collection(mongo.CollectionFoundationProposal).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		common.Logger.Errorf("change proposal status error. err: %s", err)
		return err
	}

	proposal, err := foundationProposalService.GetProposal(proposalId)
	if err != nil {
		return err
	}

	err = rollbackService.Insert(mongo.CollectionFoundationProposal, proposal.ID, height, "status", proposal.Status, status)
	if err != nil {
		common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionFoundationProposal, err)
		return err
	}

	return nil
}

func (foundationProposalService FoundationProposalService) GetProposal(proposalId int64) (*model.FoundationProposal, error) {
	var result model.FoundationProposal
	filter := bson.M{"proposal_id": proposalId}
	err := mongo.FindOne(mongo.CollectionFoundationProposal, filter, &result)
	return &result, err
}
