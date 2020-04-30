package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
)

type FoundationMemberService struct{}

func (foundationMemberService FoundationMemberService) Insert(height int64, time int64, members []string) error {
	inserts := make([]interface{}, 0)

	for _, v := range members {
		m := model.FoundationMember{
			Time:      time,
			Address:   v,
			Height:    height,
			InService: true,
		}
		inserts = append(inserts, m)
	}

	_, err := mongo.MongoDB.Collection(mongo.CollectionFoundationMember).InsertMany(context.TODO(), inserts)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionFoundationMember, err)
		return err
	}
	return nil
}

func (foundationMemberService FoundationMemberService) OutService(height int64, time int64, member string) error {
	old, err := foundationMemberService.FindInServiceMember(member)
	if err != nil {
		return err
	}

	updateFilter := bson.M{"_id": old.ID}
	_, err = mongo.MongoDB.Collection(mongo.CollectionFoundationMember).UpdateOne(context.TODO(), updateFilter, bson.M{"$set": bson.M{"in_service": false}})
	if err != nil {
		return nil
	}
	err = rollbackService.Insert(mongo.CollectionFoundationMember, old.ID, height, "in_service", old.InService, false)
	if err != nil {
		common.Logger.Errorf("insert rollback error. rollback collection: %s, err: %s", mongo.CollectionFoundationMember, err)
		return err
	}

	return nil
}

func (foundationMemberService FoundationMemberService) FindInServiceMember(address string) (*model.FoundationMember, error) {
	var result model.FoundationMember
	filter := bson.M{"address": address, "in_service": true}
	err := mongo.FindOne(mongo.CollectionFoundationMember, filter, &result)
	return &result, err
}
