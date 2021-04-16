package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
)

type ConfigService struct{}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (s *ConfigService) InfoConfig(config *model.Config) {
	if !config.Info() {
		panic(NewRespErr(ErrNotExist, "不存在的配置记录"))
	}
	return
}

func (s *ConfigService) ListConfig(args dao.ListConfigArgs) MapItf {
	data := MapItf{}
	list := dao.ListConfig(args)
	for _, v := range list {
		data[v.Key] = v
	}
	return data
}
