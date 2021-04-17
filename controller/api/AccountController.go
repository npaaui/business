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
		"user_bank_id": "int|required",
		"amount":       "float|required",
		"img":          "string|required",
	}, accountInOut)
	c.service.Recharge(accountInOut)
	ReturnData(g, accountInOut)
}

func (c *AccountController) Withdraw(g *gin.Context) {
	args := &service.WithdrawArgs{
		UserId: TokenInfo.UserId,
	}
	ValidatePostJson(g, map[string]string{
		"user_bank_id": "int|required||用户银行卡编号",
		"amount":       "float|required||提现金额",
		"password":     "string|required||提现密码",
	}, args)

	c.service.Withdraw(args)
	ReturnData(g, nil)
}
