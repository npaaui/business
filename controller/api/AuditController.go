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
	args := &dao.ListAuditArgs{
		UserId: TokenInfo.UserId,
	}
	ValidateQuery(g, map[string]string{
		"action":            "string",
		"status":            "string",
		"create_time_start": "string",
		"create_time_end":   "string",
		"page":              "int",
		"page_size":         "int",
	}, args)
	list := c.service.ListAudit(args)
	ReturnData(g, list)
}

func (c *AuditController) UpdateAudit(g *gin.Context) {
	args := &model.Audit{
		OpsId: TokenInfo.UserId,
	}
	ValidatePostJson(g, map[string]string{
		"id":     "int",
		"status": "string",
		"remark": "string",
	}, args)
	c.service.UpdateAudit(args)
	ReturnData(g, args)
}
