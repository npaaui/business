package main

import (
	"fmt"

	"github.com/npaaui/helper-go/db"
	"github.com/npaaui/helper-go/gen"

	. "business/common"
)

func main() {
	InitConfig("config.ini")
	mysqlConf, err := Conf.GetSection("MYSQL")
	if err != nil {
		panic(fmt.Errorf("mysql get conf error: %w", err))
	}
	(&db.Conf{
		DriverName:      "mysql",
		ConnMaxLifetime: 86400,
		Prefix:          mysqlConf["prefix"],
		Conn: db.MysqlConf{
			Host:     mysqlConf["host"],
			Username: mysqlConf["username"],
			Password: mysqlConf["password"],
			Database: mysqlConf["database"],
		},
	}).InitDbEngine()

	(&gen.Conf{
		ModelFolder: "dao/model/",
		TplFile:     "tmp/gen/model.tpl",
		TableNames:  "",
		DbName:      "business",
	}).InitGenConf()
	gen.GenerateModelFile()
}
