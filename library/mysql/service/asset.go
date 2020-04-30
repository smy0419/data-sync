package service

import (
	"github.com/AsimovNetwork/data-sync/library/mysql"
	models "github.com/AsimovNetwork/data-sync/library/mysql/model"
)

type AssetService struct{}

func GetAsset(asset string) (map[string]string, error) {
	assetModel := new(models.TDaoAsset)
	total, err := mysql.Engine.Where("asset = ?", asset).Count(assetModel)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return nil, nil
	}

	assetModel = new(models.TDaoAsset)
	_, err = mysql.Engine.Where("asset = ?", asset).Get(assetModel)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	result["asset"] = assetModel.Asset
	result["name"] = assetModel.Name
	result["symbol"] = assetModel.Symbol
	result["description"] = assetModel.Description
	result["logo"] = assetModel.Logo
	return result, err
}
