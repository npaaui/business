package model

import (
	"business/common"
	"fmt"
	"github.com/go-xorm/xorm"
	dbHelper "github.com/npaaui/go-helper-db"
)

var (
	DbEngine *xorm.Engine
)

func InitMysql() {
	mysqlConf, err := common.Conf.GetSection("MYSQL")
	if err != nil {
		panic(fmt.Errorf("mysql get conf error: %w", err))
	}
	dbConf := dbHelper.DbConf{
		DriverName:      "mysql",
		ConnMaxLifetime: 86400,
		Prefix:          mysqlConf["prefix"],
		Conn: dbHelper.MysqlConf{
			Host:     mysqlConf["host"],
			Username: mysqlConf["username"],
			Password: mysqlConf["password"],
			Database: mysqlConf["database"],
		},
	}
	dbConf.InitDbEngine()
	DbEngine = dbHelper.GetDbEngineIns()
}