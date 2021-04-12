package api

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
	"business/service"
	"github.com/gin-gonic/gin"
	"strings"
)

type ConfigController struct {
	service *service.ConfigService
}

func NewConfigController() *ConfigController {
	return &ConfigController{
		service: service.NewConfigService(),
	}
}

/**
 * 获取配置详情
 */
func (c *ConfigController) InfoConfig(g *gin.Context) {
	config := model.NewConfigModel()
	ValidateParam(g, map[string]string{
		"key": "string",
	}, map[string]string{
		"key": "required|string",
	}, config)

	c.service.InfoConfig(config)

	ReturnData(g, config)
}

/**
 * 获取配置列表
 */
type ListConfig struct {
	Keys string
}

func (c *ConfigController) ListConfig(g *gin.Context) {
	args := &ListConfig{}
	ValidateQuery(g, map[string]string{
		"keys": "string",
	}, map[string]string{
		"keys": "required|string",
	}, args)

	keys := strings.Split(args.Keys, ",")
	configList := c.service.ListConfig(dao.ListConfigArgs{
		Keys: keys,
	})

	ReturnData(g, configList)
}
