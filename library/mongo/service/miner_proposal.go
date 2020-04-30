package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type MinerProposalService struct{}

var rollbackService = RollbackService{}

func (minerProposalService MinerProposalService) Insert(
	round int64,
	height int64,
	time int64,
	endTime int64,
	proposalId int64,
	proposer string,
	proposalType uint8,
	status uint8,
	txHash string,
) error {
	proposal := model.MinerProposal{
		Round:           round,
		Height:          height,
		Time:            time,
		EndTime:         endTime,
		ProposalId:      proposalId,
		Address:         proposer,
		Type:            proposalType,
		Status:          status,
		TxHash:          txHash,
		EffectiveHeight: 0,
		EffectiveTime:   0,
		SupportRate:     -1,
		RejectRate:      -1,
	}

	_, err := mongo.MongoDB.Collection(mongo.CollectionMinerProposal).InsertOne(context.TODO(), proposal)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionMinerProposal, err)
		return err
	}

	return nil
}

func (minerProposalService MinerProposalService) GetProposal(proposalId int64) (*model.MinerProposal, error) {
	var result model.MinerProposal
	filter := bson.M{"proposal_id": proposalId}
	err := mongo.FindOne(mongo.CollectionMinerProposal, filter, &result)
	return &result, err
}

func (minerProposalService MinerProposalService) ChangeStatus(height int64, time int64, round int64, proposalId int64, status uint8, supportRate int64, rejectRate int64) error {
	proposal, err := minerProposalService.GetProposal(proposalId)
	if err != nil {
		return err
	}

	filter := bson.M{"proposal_id": proposalId}
	update := bson.M{
		"status":       status,
		"support_rate": supportRate,
		"reject_rate":  rejectRate,
	}
	_, err = mongo.MongoDB.Collection(mongo.CollectionMinerProposal).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		common.Logger.Errorf("change miner proposal status error. err: %s", err)
		return err
	}

	err = rollbackService.Insert(mongo.CollectionMinerProposal, proposal.ID, height, "status", proposal.Status, status)
	if err != nil {
		common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionMinerProposal, err)
		return err
	}

	err = rollbackService.Insert(mongo.CollectionMinerProposal, proposal.ID, height, "support_rate", proposal.SupportRate, supportRate)
	if err != nil {
		common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionMinerProposal, err)
		return err
	}

	err = rollbackService.Insert(mongo.CollectionMinerProposal, proposal.ID, height, "reject_rate", proposal.RejectRate, rejectRate)
	if err != nil {
		common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionMinerProposal, err)
		return err
	}

	if status == model.MinerProposalStatusReject || status == model.MinerProposalStatusEffective {
		err = minerTodoListService.Release(height, time, round, proposal.Address, proposalId)
		if err != nil {
			common.Logger.Errorf("release miner to do list error. err: %s", err)
			return err
		}
	}

	return nil
}

func (minerProposalService MinerProposalService) ChangeStatusByWorkHeight(height int64, proposalId int64, workHeight int64) error {
	proposal, err := minerProposalService.GetProposal(proposalId)
	if err != nil {
		return err
	}

	filter := bson.M{"proposal_id": proposalId}

	effectiveTime := common.GetEffectiveTime(height, workHeight)
	var update interface{}
	if workHeight == 0 {
		update = bson.M{"status": model.MinerProposalStatusEffective}
	} else {
		update = bson.M{"effective_height": workHeight, "effective_time": effectiveTime}
	}

	_, err = mongo.MongoDB.Collection(mongo.CollectionMinerProposal).UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		common.Logger.Errorf("change miner proposal status error. err: %s", err)
		return err
	}

	if workHeight == 0 {
		err = rollbackService.Insert(mongo.CollectionMinerProposal, proposal.ID, height, "status", proposal.Status, model.MinerProposalStatusEffective)
	} else {
		err = rollbackService.Insert(mongo.CollectionMinerProposal, proposal.ID, height, "effective_height", proposal.EffectiveHeight, workHeight)
		err = rollbackService.Insert(mongo.CollectionMinerProposal, proposal.ID, height, "effective_time", proposal.EffectiveTime, effectiveTime)
	}
	if err != nil {
		common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionMinerProposal, err)
		return err
	}

	return nil
}

func (minerProposalService MinerProposalService) UpdateStatusByHeight(height int64) error {
	filter := bson.M{
		"effective_height": bson.M{
			"$lte": height,
		},
		"status": model.MinerProposalStatusApproved,
	}
	update := bson.M{
		"status": model.MinerProposalStatusEffective,
	}
	_, err := mongo.MongoDB.Collection(mongo.CollectionMinerProposal).UpdateMany(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		common.Logger.Errorf("change miner proposal status error. err: %s", err)
		return err
	}
	proposals, err := mongo.Find(mongo.CollectionMinerProposal, filter, reflect.TypeOf(model.Rollback{}), reflect.TypeOf(&model.MinerProposal{}))
	for _, v := range proposals.([]*model.MinerProposal) {
		err = rollbackService.Insert(mongo.CollectionMinerProposal, v.ID, height, "status", v.Status, model.MinerProposalStatusEffective)
		if err != nil {
			common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionMinerProposal, err)
			return err
		}
	}

	return nil
}
