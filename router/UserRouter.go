package router

import (
	"github.com/gin-gonic/gin"

	"business/controller/api"
)

func LoadUserRouter(r gin.IRoutes) {
	userCtrl := api.NewUserController()
	shopCtrl := api.NewShopController()
	userBankCtrl := api.NewUserBankController()
	accountCtrl := api.NewAccountController()
	taskCtrl := api.NewTaskController()

	u := r
	{
		u.GET("/user", userCtrl.InfoUser)
		u.GET("/user/list", userCtrl.ListUser)
		u.PUT("/user/password", userCtrl.ResetUserPassword)

		u.GET("/shop", shopCtrl.ListShop)
		u.POST("/shop", shopCtrl.InsertShop)
		u.PUT("/shop", shopCtrl.UpdateShop)

		u.GET("/user_bank", userBankCtrl.ListUserBank)
		u.POST("/user_bank", userBankCtrl.InsertUserBank)
		u.PUT("/user_bank", userBankCtrl.UpdateUserBank)
		u.DELETE("/user_bank", userBankCtrl.DeleteUserBank)

		u.POST("/recharge", accountCtrl.Recharge)
		u.POST("/withdraw", accountCtrl.Withdraw)
		u.GET("/account_in_out", accountCtrl.ListAccountInOut)
		u.PUT("/account_in_out/status", accountCtrl.UpdateAccountInOutStatus)
		u.GET("/account_log", accountCtrl.ListAccountLog)

		u.GET("/task", taskCtrl.ListTask)
		u.GET("/task/:id", taskCtrl.InfoTask)
		u.POST("/task", taskCtrl.InsertTask)
		u.PUT("/task/status", taskCtrl.UpdateTaskStatus)
		u.GET("/order", taskCtrl.ListOrder)
	}
}
