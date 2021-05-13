package router

import (
	"github.com/gin-gonic/gin"

	"business/controller/api"
)

func LoadAuditRouter(r gin.IRoutes) {
	auditCtrl := api.NewAuditController()

	a := r
	{
		a.GET("/audit", auditCtrl.ListAudit)
		a.PUT("/audit", auditCtrl.UpdateAudit)
	}
}
