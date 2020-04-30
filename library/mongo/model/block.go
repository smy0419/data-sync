package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Block struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Hash              string             `bson:"hash" json:"hash"`
	Confirmations     int64              `bson:"confirmations" json:"confirmations"`
	Size              int32              `bson:"size" json:"size"`
	Height            int64              `bson:"height" json:"height"`
	Version           int32              `bson:"version" json:"version"`
	MerkleRoot        string             `bson:"merkle_root" json:"merkle_root"`
	Time              int64              `bson:"time" json:"time"`
	TxCount           uint64             `bson:"tx_count" json:"tx_count"`
	PreviousBlockHash string             `bson:"previous_block_hash" json:"previous_block_hash"`
	StateRoot         string             `bson:"state_root" json:"state_root"`
	// block produced address
	Produced string `bson:"produced" json:"produced"`
	Reward   int64  `bson:"reward" json:"reward"`
	Fee      []Fee  `bson:"fee" json:"fee"`
	Slot     int32  `bson:"slot" json:"slot"`
	Round    int32  `bson:"round" json:"round"`
}
