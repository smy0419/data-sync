package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	MinerProposalStatusOnGoing uint8 = iota
	MinerProposalStatusApproved
	MinerProposalStatusReject
	MinerProposalStatusEffective
)

type MinerProposal struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Round           int64              `bson:"round" json:"round"`
	Height          int64              `bson:"height" json:"height"`
	Time            int64              `bson:"time" json:"time"`
	EndTime         int64              `bson:"end_time" json:"end_time"`
	ProposalId      int64              `bson:"proposal_id" json:"proposal_id"`
	Address         string             `bson:"address" json:"address"`
	Type            uint8              `bson:"type" json:"type"`
	Status          uint8              `bson:"status" json:"status"`
	TxHash          string             `bson:"tx_hash" json:"tx_hash"`
	EffectiveHeight int64              `bson:"effective_height" json:"effective_height"`
	EffectiveTime   int64              `bson:"effective_time" json:"effective_time"`
	SupportRate     int64              `bson:"support_rate" json:"support_rate"`
	RejectRate      int64              `bson:"reject_rate" json:"reject_rate"`
}

type MinerVote struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Round      int64              `bson:"round" json:"round"`
	Height     int64              `bson:"height" json:"height"`
	Time       int64              `bson:"time" json:"time"`
	ProposalId int64              `bson:"proposal_id" json:"proposal_id"`
	Voter      string             `bson:"voter" json:"voter"`
	Decision   bool               `bson:"decision" json:"decision"`
	TxHash     string             `bson:"tx_hash" json:"tx_hash"`
}

type MinerMember struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height     int64              `bson:"height" json:"height"`
	Time       int64              `bson:"time" json:"time"`
	Round      int64              `bson:"round" json:"round"`
	Address    string             `bson:"address" json:"address"`
	Produced   int64              `bson:"produced" json:"produced"`
	Planed     int64              `bson:"planed" json:"planed"`
	Efficiency int32              `bson:"efficiency" json:"efficiency"`
}

type MinerRound struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height    int64              `bson:"height" json:"height"`
	Time      int64              `bson:"time" json:"time"`
	Round     int64              `bson:"round" json:"round"`
	StartTime int64              `bson:"start_time" json:"start_time"`
	EndTime   int64              `bson:"end_time" json:"end_time"`
}

type MinerTodoList struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height     int64              `bson:"height" json:"height"`
	Time       int64              `bson:"time" json:"time"`
	Round      int64              `bson:"round" json:"round"`
	Operator   string             `bson:"operator" json:"operator"`
	ActionId   int64              `bson:"action_id" json:"action_id"`
	ActionType uint8              `bson:"action_type" json:"action_type"`
	EndTime    int64              `bson:"end_time" json:"end_time"`
	Operated   bool               `bson:"operated" json:"operated"`
}

type MinerSignUp struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height     int64              `bson:"height" json:"height"`
	Time       int64              `bson:"time" json:"time"`
	TxHash     string             `bson:"tx_hash" json:"tx_hash"`
	Round      int64              `bson:"round" json:"round"`
	Address    string             `bson:"address" json:"address"`
	Produced   int64              `bson:"produced" json:"produced"`
	Planed     int64              `bson:"planed" json:"planed"`
	Efficiency int32              `bson:"efficiency" json:"efficiency"`
}
