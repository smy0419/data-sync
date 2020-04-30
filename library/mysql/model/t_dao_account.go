package models

type TDaoAccount struct {
	Id         int64  `xorm:"pk comment('Primary Key') BIGINT(64)"`
	Address    string `xorm:"not null comment('Address') unique VARCHAR(80)"`
	NickName   string `xorm:"not null comment('Nick Name') VARCHAR(64)"`
	Avatar     string `xorm:"not null default '' comment('Avatar') VARCHAR(256)"`
	CreateTime int64  `xorm:"not null comment('Create Time') BIGINT(20)"`
	UpdateTime int64  `xorm:"not null comment('Update Time') BIGINT(20)"`
}
