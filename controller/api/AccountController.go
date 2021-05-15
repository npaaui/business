package api

import (
	"github.com/gin-gonic/gin"
	"strings"

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
		SetUserId(g.GetInt("user_id")).
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
		UserId: g.GetInt("user_id"),
	}
	ValidatePostJson(g, map[string]string{
		"user_bank_id": "int|required||用户银行卡编号",
		"amount":       "float|required||提现金额",
		"password":     "string|required||提现密码",
	}, args)

	c.service.Withdraw(args)
	ReturnData(g, nil)
}

func (c *AccountController) UpdateAccountInOutStatus(g *gin.Context) {
	args := &model.AccountInOut{
		UserId: g.GetInt("user_id"),
	}
	ValidatePostJson(g, map[string]string{
		"id":     "int|required",
		"status": "string|required|enum:" + dao.AccountInOutStatusCancel,
	}, args)

	c.service.UpdateAccountInOut(args)
	ReturnData(g, nil)
}

/**
 * 获取充提申请记录列表
 */
func (c *AccountController) ListAccountInOut(g *gin.Context) {
	args := &dao.ListAccountInOutArgs{
		UserId: g.GetInt("user_id"),
	}
	ValidateQuery(g, map[string]string{
		"type":      "string|enum:" + strings.Join(dao.AccountInOutTypeSlice, ","),
		"page":      "int",
		"page_size": "int",
	}, args)
	accountInOutList := c.service.ListAccountInOut(args)
	ReturnData(g, accountInOutList)
}

/**
 * 获取资金记录列表
 */
func (c *AccountController) ListAccountLog(g *gin.Context) {
	args := &dao.ListAccountLogArgs{
		UserId: g.GetInt("user_id"),
	}
	ValidateQuery(g, map[string]string{
		"account_type":      "string|enum:" + strings.Join(dao.AccountTypeSlice, ","),
		"type":              "string|enum:" + strings.Join(dao.AccountLogTypeSlice, ","),
		"task_id":           "int",
		"order_id":          "int",
		"shop_id":           "int",
		"create_time_start": "string",
		"create_time_end":   "string",
		"page":              "int",
		"page_size":         "int",
	}, args)
	accountInOutList := c.service.ListAccountLog(args)
	ReturnData(g, accountInOutList)
}
