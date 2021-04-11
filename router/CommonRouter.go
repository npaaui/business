package router

import (
	"github.com/gin-gonic/gin"

	"business/controller/api"
)

func LoadCommonRouter(r gin.IRoutes) {
	categoryCtrl := api.NewCategoryController()

	c := r
	{
		c.GET("/category", categoryCtrl.ListCategory)
	}
}
