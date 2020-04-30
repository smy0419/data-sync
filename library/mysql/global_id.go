package mysql

import "github.com/AsimovNetwork/data-sync/library/common"

type GlobalId struct {
	Snowflake *common.Snowflake
}

func (globalIdService GlobalId) NextId() int64 {
	return globalIdService.Snowflake.Generate()
}
