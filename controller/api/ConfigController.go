package api

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
	"business/service"
	"github.com/gin-gonic/gin"
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
func (c *ConfigController) ListConfig(g *gin.Context) {
	keys := g.QueryArray("keys")
	if len(keys) == 0 {
		ReturnErrMsg(g, ErrValidReq, "配置Keys值不可为空")
		return
	}

	configList := c.service.ListConfig(dao.ListConfigArgs{
		Keys: keys,
	})

	ReturnData(g, configList)
}
