package router

import (
	"github.com/gin-gonic/gin"

	"business/controller/api"
)

func LoadCommonRouter(r gin.IRoutes) {
	uploadCtrl := api.NewUploadController()
	categoryCtrl := api.NewCategoryController()
	configCtrl := api.NewConfigController()

	c := r
	{
		c.POST("/file", uploadCtrl.UploadFile)
		c.GET("/category", categoryCtrl.ListCategory)
		c.GET("/config/:key", configCtrl.InfoConfig)
		c.GET("/config", configCtrl.ListConfig)
	}
}
