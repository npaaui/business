package common

import (
	"fmt"

	"github.com/Unknwon/goconfig"
)

var (
	Conf *goconfig.ConfigFile
	ConfCom = &ConfCommon{}
)

type ConfCommon struct {
	Env 	string
}

func InitConfig(path string) {
	var err error
	Conf, err = goconfig.LoadConfigFile(path)
	if err != nil {
		panic(fmt.Errorf("config init error: %w", err))
	}

	commonConf, err := Conf.GetSection("COMMON")
	if err != nil {
		panic(fmt.Errorf("init common conf error: %w", err))
	}
	ConfCom.Env = commonConf["env"]
}

