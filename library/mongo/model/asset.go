package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CreateAsset = 1 // create
	IssueAsset  = 2 // mint
)

type Asset struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height       int64              `bson:"height" json:"height"`
	Time         int64              `bson:"time" json:"time"`
	IssueAddress string             `bson:"issue_address" json:"issue_address"`
	Asset        string             `bson:"asset" json:"asset"`
	Name         string             `bson:"name" json:"name"`
	Symbol       string             `bson:"symbol" json:"symbol"`
	Description  string             `bson:"description" json:"description"`
	Logo         string             `bson:"logo" json:"logo"`
}

type AssetIssue struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Asset     string             `bson:"asset" json:"asset"`
	Height    int64              `bson:"height" json:"height"`
	Time      int64              `bson:"time" json:"time"`
	IssueType uint8              `bson:"issue_type" json:"issue_type"`
	// issue value
	Value int64 `bson:"value" json:"value"`
}
