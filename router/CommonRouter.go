package router

import (
	"github.com/gin-gonic/gin"

	"business/controller/api"
)

func LoadCommonRouter(r gin.IRoutes) {
	categoryCtrl := api.NewCategoryController()
	configCtrl := api.NewConfigController()

	c := r
	{
		c.GET("/category", categoryCtrl.ListCategory)
		c.GET("/config/:key", configCtrl.InfoConfig)
		c.GET("/config", configCtrl.ListConfig)
	}
}
