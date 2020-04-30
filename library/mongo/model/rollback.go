package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rollback struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	// mongodb collection name
	Collection string `bson:"collection" json:"collection"`
	// object id
	DocumentID primitive.ObjectID `bson:"document_id" json:"document_id"`
	// block height
	Height int64 `bson:"height" json:"height"`
	// record creation time
	Time int64 `bson:"time" json:"time"`
	// rollback field
	Field         string      `bson:"field" json:"field"`
	OriginalValue interface{} `bson:"original_value" json:"original_value"`
	ExpectValue   interface{} `bson:"expect_value" json:"expect_value"`
}
