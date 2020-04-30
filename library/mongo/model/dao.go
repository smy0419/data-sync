package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	OrgStatusNormal uint8 = iota
	OrgStatusClosed
)

const (
	MemberStatusInvited uint8 = iota
	MemberStatusAgreed
	MemberStatusRemoved
)

const (
	MemberRolePresident uint8 = iota
	MemberRoleOrdinary
)

const (
	ProposalTypeIssueAsset uint8 = iota
)

const (
	TodoTypeInviteMember uint8 = iota
	TodoTypeInvitePresident
	TodoTypeVote
)

const (
	AssetStatusInit uint8 = iota
	AssetStatusSuccess
)

const (
	AssetTypeDivisible uint32 = iota
	AssetTypeIndivisible
)

type DaoOrganization struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height              int64              `bson:"height" json:"height"`
	Time                int64              `bson:"time" json:"time"`
	TxHash              string             `bson:"tx_hash" json:"tx_hash"`
	OrgId               uint32             `bson:"org_id" json:"org_id"`
	ContractAddress     string             `bson:"contract_address" json:"contract_address"`
	VoteContractAddress string             `bson:"vote_contract_address" json:"vote_contract_address"`
	VoteTemplateName    string             `bson:"vote_template_name" json:"vote_template_name"`
	// president address
	President string `bson:"president" json:"president"`
	OrgName   string `bson:"org_name" json:"org_name"`
	Status    uint8  `bson:"status" json:"status"`
}

type DaoOrganizationAsset struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height          int64              `bson:"height" json:"height"`
	Time            int64              `bson:"time" json:"time"`
	ContractAddress string             `bson:"contract_address" json:"contract_address"`
	Asset           string             `bson:"asset" json:"asset"`
	AssetType       uint32             `bson:"asset_type" json:"asset_type"`
	AssetIndex      uint32             `bson:"asset_index" json:"asset_index"`
}

type DaoProposal struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height          int64              `bson:"height" json:"height"`
	Time            int64              `bson:"time" json:"time"`
	TxHash          string             `bson:"tx_hash" json:"tx_hash"`
	EndTime         int64              `bson:"end_time" json:"end_time"`
	ContractAddress string             `bson:"contract_address" json:"contract_address"`
	ProposalId      int64              `bson:"proposal_id" json:"proposal_id"`
	Address         string             `bson:"address" json:"address"`
	ProposalType    uint8              `bson:"proposal_type" json:"proposal_type"`
	// Proposal Status: 0-ongoing, 1-passed, 2-not passed
	Status uint8 `bson:"status" json:"status"`
}

type DaoMember struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height          int64              `bson:"height" json:"height"`
	Time            int64              `bson:"time" json:"time"`
	ContractAddress string             `bson:"contract_address" json:"contract_address"`
	Role            uint8              `bson:"role" json:"role"`
	Address         string             `bson:"address" json:"address"`
	Status          uint8              `bson:"status" json:"status"`
}

type DaoTodoList struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height          int64              `bson:"height" json:"height"`
	Time            int64              `bson:"time" json:"time"`
	ContractAddress string             `bson:"contract_address" json:"contract_address"`
	Operator        string             `bson:"operator" json:"operator"`
	TodoType        uint8              `bson:"todo_type" json:"todo_type"`
	TodoId          int64              `bson:"todo_id" json:"todo_id"`
	EndTime         int64              `bson:"end_time" json:"end_time"`
	Operated        bool               `bson:"operated" json:"operated"`
}

type DaoVote struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height          int64              `bson:"height" json:"height"`
	Time            int64              `bson:"time" json:"time"`
	TxHash          string             `bson:"tx_hash" json:"tx_hash"`
	ContractAddress string             `bson:"contract_address" json:"contract_address"`
	Voter           string             `bson:"voter" json:"voter"`
	VoteId          int64              `bson:"vote_id" json:"vote_id"`
	Decision        bool               `bson:"decision" json:"decision"`
}
