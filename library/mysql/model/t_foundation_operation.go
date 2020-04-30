package models

type TFoundationOperation struct {
	Id             int64  `xorm:"pk comment('Primary Key') BIGINT(64)"`
	TxHash         string `xorm:"not null default '' comment('Transaction Hash') unique VARCHAR(64)"`
	OperationType  int    `xorm:"not null comment('Operate Type: 1.proposal 2.vote 3.donate') index index(INDEX_OPERATOR_TYPE) TINYINT(4)"`
	AdditionalInfo string `xorm:"not null comment('Additional Information(JSON Format)') VARCHAR(1024)"`
	TxStatus       int    `xorm:"not null default 0 comment('Transaction Status: 0.unconfirmedï¼Œ1.transaction confirmed, contract execution success 2.transaction confirmed, contract execution failed') TINYINT(4)"`
	Operator       string `xorm:"not null comment('Operator Address') index(INDEX_OPERATOR_TYPE) VARCHAR(64)"`
	CreateTime     int64  `xorm:"not null comment('Create Time') BIGINT(20)"`
	UpdateTime     int64  `xorm:"not null comment('Update Time') BIGINT(20)"`
}
