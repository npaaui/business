package api

import (
	"github.com/gin-gonic/gin"

	. "business/common"
	"business/dao"
	"business/dao/model"
	"business/service"
)

type AccountController struct {
	service *service.UserService
}

func NewAccountController() *AccountController {
	return &AccountController{
		service: service.NewUserService(),
	}
}

func (c *AccountController) Recharge(g *gin.Context) {
	accountInOut := model.NewAccountInOutModel().
		SetUserId(TokenInfo.UserId).
		SetType(dao.AccountInOutTypeRecharge)
	ValidatePostJson(g, map[string]string{
		"user_bank_id": "int",
		"amount":       "float",
		"img":          "string",
	}, map[string]string{
		"user_bank_id": "required|int",
		"amount":       "required|float",
		"img":          "required|string",
	}, accountInOut)
	c.service.Recharge(accountInOut)
	ReturnData(g, accountInOut)
}

func (c *AccountController) Withdraw(g *gin.Context) {
	args := &service.WithdrawArgs{
		UserId: TokenInfo.UserId,
	}
	ValidatePostJson(g, map[string]string{
		"user_bank_id": "int",
		"amount":       "float",
		"password":     "string",
	}, map[string]string{
		"user_bank_id": "required|int",
		"amount":       "required|float",
		"password":     "required|string",
	}, args)

	c.service.Withdraw(args)
	ReturnData(g, nil)
}
