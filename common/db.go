package common

import (
	"fmt"
	"github.com/npaaui/helper-go/tools"

	"github.com/go-xorm/xorm"
	"github.com/npaaui/helper-go/db"
)

var (
	DbEngine       *xorm.Engine
	UniqueIdWorker *tools.Worker
)

// 初始化数据库
func InitMysql() {
	mysqlConf, err := Conf.GetSection("MYSQL")
	if err != nil {
		panic(fmt.Errorf("get mysql conf error: %w", err))
	}
	dbConf := db.Conf{
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
	if mysqlConf["show_sql"] == "true" {
		DbEngine.ShowSQL(true)
	}
}
