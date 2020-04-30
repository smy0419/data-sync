package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContractTemplate struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Height    int64              `bson:"height" json:"height"`
	Time      int64              `bson:"time" json:"time"`
	Owner     string             `bson:"owner" json:"owner"`
	Category  uint16             `bson:"category" json:"category"`
	Name      string             `bson:"name" json:"name"`
	Key       string             `bson:"key" json:"key"`
	Approver  uint8              `bson:"approver" json:"approver"`
	Rejecter  uint8              `bson:"rejecter" json:"rejecter"`
	Committee uint8              `bson:"committee" json:"committee"`
	Status    uint8              `bson:"status" json:"status"`
}
