package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ProposalStatusOnGoing uint8 = iota
	ProposalStatusApproved
	ProposalStatusReject
)

const (
	TransferTypeIn uint8 = iota
	TransferTypeOut
)

type FoundationProposal struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height       int64              `bson:"height" json:"height"`
	Time         int64              `bson:"time" json:"time"`
	EndTime      int64              `bson:"end_time" json:"end_time"`
	ProposalId   int64              `bson:"proposal_id" json:"proposal_id"`
	Address      string             `bson:"address" json:"address"`
	ProposalType uint8              `bson:"proposal_type" json:"proposal_type"`
	Status       uint8              `bson:"status" json:"status"`
	TxHash       string             `bson:"tx_hash" json:"tx_hash"`
}

type FoundationVote struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height     int64              `bson:"height" json:"height"`
	Time       int64              `bson:"time" json:"time"`
	ProposalId int64              `bson:"proposal_id" json:"proposal_id"`
	Voter      string             `bson:"voter" json:"voter"`
	Decision   bool               `bson:"decision" json:"decision"`
	TxHash     string             `bson:"tx_hash" json:"tx_hash"`
}

type FoundationBalanceSheet struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height       int64              `bson:"height" json:"height"`
	Time         int64              `bson:"time" json:"time"`
	Address      string             `bson:"address" json:"address"`
	Asset        string             `bson:"asset" json:"asset"`
	Amount       int64              `bson:"amount" json:"amount"`
	TransferType uint8              `bson:"transfer_type" json:"transfer_type"`
	ProposalType int8               `bson:"proposal_type" json:"proposal_type"`
	TxHash       string             `bson:"tx_hash" json:"tx_hash"`
}

type FoundationMember struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height    int64              `bson:"height" json:"height"`
	Time      int64              `bson:"time" json:"time"`
	Address   string             `bson:"address" json:"address"`
	InService bool               `bson:"in_service" json:"in_service"`
}

type FoundationTodoList struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height       int64              `bson:"height" json:"height"`
	Time         int64              `bson:"time" json:"time"`
	Operator     string             `bson:"operator" json:"operator"`
	TodoId       int64              `bson:"todo_id" json:"todo_id"`
	ProposalType uint8              `bson:"proposal_type" json:"proposal_type"`
	EndTime      int64              `bson:"end_time" json:"end_time"`
	Operated     bool               `bson:"operated" json:"operated"`
}
