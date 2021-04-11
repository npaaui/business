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
 * 获取任务列表
 */
func (c *CategoryController) ListCategory(g *gin.Context) {
	category := model.NewCategoryModel().SetType("sell")
	categoryList := c.service.ListCategory(*category)
	ReturnData(g, categoryList)
}
