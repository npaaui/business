package router

import (
	"github.com/gin-gonic/gin"

	"business/controller/api"
)

func LoadUserRouter(r gin.IRoutes) {
	userCtrl := api.NewUserController()
	shopCtrl := api.NewShopController()
	taskCtrl := api.NewTaskController()

	u := r
	{
		u.GET("/user", userCtrl.InfoUser)
		u.GET("/shop", shopCtrl.ListShop)
		u.POST("/shop", shopCtrl.InsertShop)
		u.PUT("/shop", shopCtrl.UpdateShop)
		u.GET("/task", taskCtrl.ListTask)
		u.GET("/task/:id", taskCtrl.InfoTask)
		u.POST("/task", taskCtrl.InsertTask)
	}
}
