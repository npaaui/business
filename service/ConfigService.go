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

type ListConfigArgs struct {
	Keys []string
}

func (s *ConfigService) ListConfig(args ListConfigArgs) MapItf {
	data := MapItf{}
	list := dao.ListConfig(dao.ListConfigArgs{
		Keys: args.Keys,
	})
	for _, v := range list {
		data[v.Key] = v
	}
	return data
}

func (s *ConfigService) ListConfigValue(args ListConfigArgs) MapStr {
	data := MapStr{}
	list := dao.ListConfig(dao.ListConfigArgs{
		Keys: args.Keys,
	})
	for _, v := range list {
		data[v.Key] = v.Value
	}
	return data
}
