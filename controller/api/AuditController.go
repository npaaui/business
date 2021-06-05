package api

import (
	"business/dao/model"
	"github.com/gin-gonic/gin"

	. "business/common"
	"business/dao"
	"business/service"
)

type AuditController struct {
	service *service.AuditService
}

func NewAuditController() *AuditController {
	return &AuditController{
		service: service.NewAuditService(),
	}
}

func (c *AuditController) ListAudit(g *gin.Context) {
	args := &dao.ListAuditArgs{}
	ValidateQuery(g, map[string]string{
		"action":            "string",
		"status":            "string",
		"create_time_start": "string",
		"create_time_end":   "string",
		"page":              "int",
		"page_size":         "int",
	}, args)

	userType := g.GetString("user_type")
	if userType != dao.UserTypeAdmin {
		args.UserId = g.GetInt("user_id")
	}

	list := c.service.ListAudit(args)
	ReturnData(g, list)
}

func (c *AuditController) UpdateAudit(g *gin.Context) {
	args := &model.Audit{
		OpsId: g.GetInt("user_id"),
	}
	ValidatePostJson(g, map[string]string{
		"id":     "string",
		"status": "string",
		"remark": "string",
	}, args)
	c.service.UpdateAudit(args)
	ReturnData(g, args)
}
