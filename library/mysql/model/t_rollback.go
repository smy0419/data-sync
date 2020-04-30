package models

type TRollback struct {
	Id            int64  `xorm:"pk comment('Primary Key') BIGINT(64)"`
	Height        int64  `xorm:"not null comment('Block Height') index BIGINT(20)"`
	RecordId      int64  `xorm:"not null comment('Roll Back Table ID') BIGINT(64)"`
	TableName     string `xorm:"not null comment('Roll Back Table Name') VARCHAR(32)"`
	OriginalValue int    `xorm:"comment('Original Value') TINYINT(4)"`
	ExpectValue   int    `xorm:"comment('Expect Value') TINYINT(4)"`
	CreateTime    int64  `xorm:"not null comment('Create Time') BIGINT(20)"`
}
