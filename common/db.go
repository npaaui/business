package common

import (
	"fmt"
	"github.com/go-xorm/xorm"

	"github.com/npaaui/go-helper/db"
)

var (
	DbEngine *xorm.Engine
)

// 初始化数据库
func InitMysql() {
	mysqlConf, err := Conf.GetSection("MYSQL")
	if err != nil {
		panic(fmt.Errorf("mysql get conf error: %w", err))
	}
	dbConf := db.DbConf{
		DriverName:      "mysql",
		ConnMaxLifetime: 86400,
		Prefix:          mysqlConf["prefix"],
		Conn: db.MysqlConf{
			Host:     mysqlConf["host"],
			Username: mysqlConf["username"],
			Password: mysqlConf["password"],
			Database: mysqlConf["database"],
		},
	}
	dbConf.InitDbEngine()
	DbEngine = db.GetDbEngineIns()
}
