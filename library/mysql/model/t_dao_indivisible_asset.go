package models

type TDaoIndivisibleAsset struct {
	Id              int64  `xorm:"pk comment('Primary Key') BIGINT(64)"`
	TxHash          string `xorm:"not null comment('Transaction Hash') VARCHAR(64)"`
	ContractAddress string `xorm:"not null comment('Contract Address') VARCHAR(64)"`
	Asset           string `xorm:"not null comment('Asset ID') VARCHAR(64)"`
	VoucherId       int64  `xorm:"not null comment('Asset Number') BIGINT(20)"`
	AssetDesc       string `xorm:"not null default '' comment('Asset Description') VARCHAR(256)"`
	AssetStatus     int    `xorm:"not null default 0 comment('Asset Status: 0.initialize 1.issue success 2.issue failed') TINYINT(4)"`
	CreateTime      int64  `xorm:"not null comment('Create Time') BIGINT(20)"`
	UpdateTime      int64  `xorm:"not null comment('Update Time') BIGINT(20)"`
}
