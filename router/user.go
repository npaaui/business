package router

import (
	"github.com/gin-gonic/gin"

	"business/controller/api"
)

func LoadUserRouter(r *gin.Engine) *gin.Engine {
	userCtrl := api.NewUserController()
	shopCtrl := api.NewShopController()

	u := r.Group("")
	{
		u.GET("/user", userCtrl.InfoUser)
		u.GET("/shop", shopCtrl.ListShop)
		u.POST("/shop", shopCtrl.InsertShop)
		u.PUT("/shop", shopCtrl.UpdateShop)
	}
	return r
}
