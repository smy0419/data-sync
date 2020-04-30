package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Validator struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height        int64              `bson:"height" json:"height"`
	Time          int64              `bson:"time" json:"time"`
	Address       string             `bson:"address" json:"address"`
	Location      Location           `bson:"location" json:"location"`
	PlannedBlocks int64              `bson:"planned_blocks" json:"planned_blocks"`
	ActualBlocks  int64              `bson:"actual_blocks" json:"actual_blocks"`
}

type BtcMiner struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height  int64              `bson:"height" json:"height"`
	Time    int64              `bson:"time" json:"time"`
	Address string             `bson:"address" json:"address"`
	Domain  string             `bson:"domain" json:"domain"`
}

// BtcMinerAddress + Released：true only one
type ValidatorRelation struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height          int64              `bson:"height" json:"height"`
	Time            int64              `bson:"time" json:"time"`
	BtcMinerAddress string             `bson:"btc_miner_address" json:"btc_miner_address"`
	Bind            bool               `bson:"bind" json:"bind"`
	Address         string             `bson:"address" json:"address"`
}

type Location struct {
	IP          string `bson:"ip" json:"ip"`
	City        string `bson:"city" json:"city"`
	Subdivision string `bson:"subdivision" json:"subdivision"`
	Country     string `bson:"country" json:"country"`
	Longitude   string `bson:"longitude" json:"longitude"`
	Latitude    string `bson:"latitude" json:"latitude"`
}

type Earning struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Time    int64              `bson:"time" json:"time"`
	Height  int64              `bson:"height" json:"height"`
	TxHash  string             `bson:"tx_hash" json:"tx_hash"`
	Address string             `bson:"address" json:"address"`
}

type EarningAsset struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Time      int64              `bson:"time" json:"time"`
	Height    int64              `bson:"height" json:"height"`
	EarningId primitive.ObjectID `bson:"earning_id" json:"earning_id"`
	Asset     string             `bson:"asset" json:"asset"`
	Value     int64              `bson:"value" json:"value"`
}
