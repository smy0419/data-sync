package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trading struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Height int64              `bson:"height" json:"height"`
	Time   int64              `bson:"time" json:"time"`
	Value  int64              `bson:"value" json:"value"`
	Asset  string             `bson:"asset" json:"asset"`
}
