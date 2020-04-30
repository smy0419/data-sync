package models

type TDaoOperation struct {
	Id              int64  `xorm:"pk comment('Primary Key') BIGINT(64)"`
	TxHash          string `xorm:"not null default '' comment('Transaction Hash') VARCHAR(64)"`
	ContractAddress string `xorm:"default '' comment('Contract Address') index VARCHAR(64)"`
	OperationType   int    `xorm:"not null comment('Operation Type: 1.create organization 2.close organization 3.modify organization name 4.remove member 5.create asset 6.transfer asset 7.president transfer 8.president confirm 9.invite member 10.join organization 11.vote 12.modify organization logo 13.mint asset') index index(INDEX_OPERATOR_TYPE) TINYINT(4)"`
	AdditionalInfo  string `xorm:"not null comment('Additional Information(JSON Format)') VARCHAR(1024)"`
	TxStatus        int    `xorm:"not null default 0 comment('Transaction Status: 0.unconfirmedï¼Œ1.transaction confirmed, contract execution success 2.transaction confirmed, contract execution failed 3.local action') TINYINT(4)"`
	Operator        string `xorm:"not null comment('Operator Address') index(INDEX_OPERATOR_TYPE) VARCHAR(64)"`
	CreateTime      int64  `xorm:"not null comment('Create Time') BIGINT(20)"`
	UpdateTime      int64  `xorm:"not null comment('Update Time') BIGINT(20)"`
}
