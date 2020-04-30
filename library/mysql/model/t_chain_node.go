package models

type TChainNode struct {
	Id          int64  `xorm:"pk comment('Primary Key') BIGINT(64)"`
	Ip          string `xorm:"not null comment('IP') index VARCHAR(32)"`
	City        string `xorm:"not null default '' comment('City') VARCHAR(32)"`
	Subdivision string `xorm:"not null default '' comment('Province') VARCHAR(32)"`
	Country     string `xorm:"not null default '' comment('County') VARCHAR(32)"`
	Longitude   string `xorm:"not null default '' comment('Longitude') VARCHAR(32)"`
	Latitude    string `xorm:"not null default '' comment('Latitude') VARCHAR(32)"`
	CreateTime  int64  `xorm:"not null comment('Create Time') BIGINT(20)"`
	UpdateTime  int64  `xorm:"not null comment('Update Time') BIGINT(20)"`
}
