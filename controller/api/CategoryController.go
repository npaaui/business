package api

import (
	"github.com/gin-gonic/gin"

	. "business/common"
	"business/dao/model"
	"business/service"
)

type CategoryController struct {
	service *service.ConfigService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		service: service.NewConfigService(),
	}
}

/**
 * 获取类别列表
 */
func (c *CategoryController) ListCategory(g *gin.Context) {
	category := model.NewCategoryModel()
	ValidateQuery(g, map[string]string{
		"type": "string|required|enum:sell,task",
	}, category)

	categoryList := c.service.ListCategory(category)
	ReturnData(g, categoryList)
}
