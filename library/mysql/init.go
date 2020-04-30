package mysql

import (
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/go-sql-driver/mysql"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
)

var Engine *xorm.Engine

var GlobalIdService = GlobalId{
	Snowflake: common.NewSnowflake(0),
}

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", common.Cfg.Mysql)
	if err != nil {
		common.Logger.ErrorPanic("database configuration error: ", err)
	}

	Engine.SetLogger(common.Logger)
	Engine.ShowSQL(common.Cfg.ShowSql)
	Engine.SetMapper(core.GonicMapper{})
	// engine.SetMaxOpenConns(5)

	if err = Engine.Ping(); err != nil {
		common.Logger.ErrorPanic("database connection error.", err)
	}
}

func NewSession() *xorm.Session {
	return Engine.NewSession()
}

func IsDuplicateError(err error) bool {
	duplicateErr, ok := err.(*mysql.MySQLError)
	if ok && duplicateErr.Number == 1062 {
		return true
	}
	return false
}
