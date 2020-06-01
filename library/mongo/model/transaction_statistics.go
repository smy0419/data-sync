package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CountAddress  uint8 = 1
	CountContract uint8 = 2
	CountAsset    uint8 = 3
)

type TransactionCount struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Key      string             `bson:"key" json:"key"`
	TxCount  int64              `bson:"tx_count" json:"tx_count"`
	Category uint8              `bson:"category" json:"category"`
	// information of contract
	Time          int64  `bson:"time,omitempty" json:"time,omitempty"`
	TxHash        string `bson:"tx_hash,omitempty" json:"tx_hash,omitempty"`
	Creator       string `bson:"creator,omitempty" json:"creator,omitempty"`
	TemplateType  uint16 `bson:"template_type,omitempty" json:"template_type,omitempty"`
	TemplateTName string `bson:"template_name,omitempty" json:"template_name,omitempty"`
}

type KeyTransaction struct {
	// Height int64              `bson:"height" json:"height"`
	TxHash string `json:"tx_hash"`
	Time   int64  `json:"time"`
	Fee    []Fee  `json:"fee"`
}
