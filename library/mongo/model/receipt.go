package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Receipt struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height            int64              `bson:"height" json:"height"`
	TxHash            string             `bson:"tx_hash" json:"tx_hash"`
	PostState         string             `bson:"root" json:"root"`
	Status            uint64             `bson:"status" json:"status"`
	CumulativeGasUsed uint64             `bson:"cumulative_gas_used" json:"cumulative_gas_used"`
	Bloom             string             `bson:"bloom" json:"bloom"`
	ContractAddress   string             `bson:"contract_address" json:"contract_address"`
	GasUsed           uint64             `bson:"gas_used" json:"gas_used"`
	Logs              []*Log             `bson:"logs" json:"logs"`
}

type Log struct {
	// Consensus fields:
	// address of the contract that generated the event
	Address string `bson:"address" json:"address"`
	// list of topics provided by the contract.
	Topics []string `bson:"topics" json:"topics"`
	// supplied by the contract, usually ABI-encoded
	Data string `bson:"data" json:"data"`

	// Derived fields. These fields are filled in by the node
	// but not secured by consensus.
	// block in which the transaction was included
	BlockNumber uint64 `bson:"block_number" json:"block_number"`
	// hash of the transaction
	TxHash string `bson:"tx_hash" json:"tx_hash"`
	// index of the transaction in the block
	TxIndex uint `bson:"tx_index" json:"tx_index"`
	// hash of the block in which the transaction was included
	BlockHash string `bson:"block_hash" json:"block_hash"`
	// index of the log in the receipt
	Index uint `bson:"log_index" json:"log_index"`

	// The Removed field is true if this log was reverted due to a chain reorganisation.
	// You must pay attention to this field if you receive logs through a filter query.
	Removed bool `bson:"removed" json:"removed"`
}
