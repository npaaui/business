package api

import (
	. "business/common"
	"business/dao/model"
	"business/service"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		service: service.NewCategoryService(),
	}
}

/**
 * 获取类别列表
 */
func (c *CategoryController) ListCategory(g *gin.Context) {
	category := model.NewCategoryModel()
	ValidateQuery(g, map[string]string{
		"type": "string",
	}, map[string]string{
		"type": "required|string|enum:sell,task",
	}, category)

	categoryList := c.service.ListCategory(category)
	ReturnData(g, categoryList)
}
