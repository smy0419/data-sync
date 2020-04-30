package models

type TDaoOrganization struct {
	Id                  int64  `xorm:"pk comment('Primary Key') BIGINT(64)"`
	TxHash              string `xorm:"not null comment('Transaction Hash') VARCHAR(64)"`
	ContractAddress     string `xorm:"not null comment('Contract Address') VARCHAR(64)"`
	VoteContractAddress string `xorm:"not null comment('Vote Contract Address') VARCHAR(64)"`
	CreatorAddress      string `xorm:"not null default '' comment('Creator Address') index VARCHAR(64)"`
	OrgName             string `xorm:"not null comment('Organization Name') VARCHAR(64)"`
	OrgLogo             string `xorm:"default '' comment('Organization Logo') VARCHAR(256)"`
	State               int    `xorm:"not null comment('Organization Status: 1.normal 2.closed 3.local initialize') TINYINT(2)"`
	CreateTime          int64  `xorm:"not null comment('Create Time') BIGINT(20)"`
	UpdateTime          int64  `xorm:"not null comment('Update Time') BIGINT(20)"`
}
