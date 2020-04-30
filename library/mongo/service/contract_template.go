package service

import (
	"context"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
)

type ContractTemplateService struct{}

func (contractTemplateService ContractTemplateService) Insert(height int64, time int64, owner string, category uint16, name string, key string, appovers uint8, rejecters uint8, committee uint8, status uint8) error {
	contractTemplate := model.ContractTemplate{
		Height:    height,
		Time:      time,
		Owner:     owner,
		Category:  category,
		Name:      name,
		Key:       key,
		Approver:  appovers,
		Rejecter:  rejecters,
		Committee: committee,
		Status:    status,
	}

	_, err := mongo.MongoDB.Collection(mongo.CollectionContractTemplate).InsertOne(context.TODO(), contractTemplate)
	if err != nil {
		common.Logger.Errorf("insert %s error. err: %s", mongo.CollectionContractTemplate, err)
		return err
	}
	return nil
}
