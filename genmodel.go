package main

import (
	"fmt"

	dbHelper "github.com/npaaui/go-helper-db"
	"github.com/npaaui/go-helper-db/gen"

	. "business/common"
)

func main() {
	InitConfig("config.ini")
	mysqlConf, err := Conf.GetSection("MYSQL")
	if err != nil {
		panic(fmt.Errorf("mysql get conf error: %w", err))
	}
	(&dbHelper.DbConf{
		DriverName:      "mysql",
		ConnMaxLifetime: 86400,
		Prefix:          mysqlConf["prefix"],
		Conn: dbHelper.MysqlConf{
			Host:     mysqlConf["host"],
			Username: mysqlConf["username"],
			Password: mysqlConf["password"],
			Database: mysqlConf["database"],
		},
	}).InitDbEngine()

	(&gen.Conf{
		ModelFolder: "tmp/gen/model/",
		TplFile:     "tmp/gen/model.tpl",
		TableNames:  "",
		DbName:      "business",
	}).InitGenConf()
	gen.GenerateModelFile()
}
