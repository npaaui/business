package api

import (
	"strings"

	"github.com/gin-gonic/gin"

	. "business/common"
	"business/dao/model"
	"business/service"
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
		"key": "string|required",
	}, config)

	c.service.InfoConfig(config)

	ReturnData(g, config)
}

/**
 * 获取配置列表
 */
type ListConfig struct {
	Keys string `json:"keys"`
}

func (c *ConfigController) ListConfig(g *gin.Context) {
	args := &ListConfig{}
	ValidateQuery(g, map[string]string{
		"keys": "string|required||配置键名",
	}, args)

	keys := strings.Split(args.Keys, ",")
	configList := c.service.ListConfig(service.ListConfigArgs{
		Keys: keys,
	})

	ReturnData(g, configList)
}
