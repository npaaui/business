package dao

import (
	. "business/common"
	"business/dao/model"
)

const (
	AuditStatusInit = "init"
	AuditStatusStop = "stop"
	AuditStatusFail = "fail"
	AuditStatusPass = "pass"
)

var AuditStatusMap = MapStr{
	AuditStatusInit: "待审核",
	AuditStatusStop: "审核中断",
	AuditStatusFail: "拒绝",
	AuditStatusPass: "通过",
}

const (
	AuditActionCodeRecharge = "recharge"
	AuditActionCodeWithdraw = "withdraw"
	AuditActionCodeTask     = "task"
)

func InsertAudit(audit *model.Audit) {
	audit.Insert()

	audit.Info()

	// 添加审核日志
	log := &model.AuditLog{
		AuditId:    audit.Id,
		Status:     audit.Status,
		LinkId:     audit.LinkId,
		Remark:     audit.Remark,
		OpsId:      audit.OpsId,
		CreateTime: GetNow(),
	}
	log.Insert()
}

/**
 * 获取审核列表
 */
type ListAuditArgs struct {
	UserId          int    `json:"user_id"`
	OpsId           int    `json:"ops_id"`
	Action          string `json:"action"`
	Status          string `json:"status"`
	CreateTimeStart string `json:"create_time_start"`
	CreateTimeEnd   string `json:"create_time_end"`
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
}

func ListAudit(args *ListAuditArgs) (int, []model.Audit) {
	session := DbEngine.Table("b_audit").Alias("a").
		Select("*").
		Where("1=1")

	if args.UserId > 0 {
		session.And("a.user_id = ?", args.UserId)
	}
	if args.OpsId > 0 {
		session.And("a.ops_id = ?", args.OpsId)
	}
	if args.Action != "" {
		session.And("a.action = ?", args.Action)
	}
	if args.Status != "" {
		session.And("a.status = ?", args.Status)
	}
	if args.CreateTimeStart != "" {
		session.And("a.create_time >= ?", args.CreateTimeStart)
	}
	if args.CreateTimeEnd != "" {
		session.And("a.create_time <= ?", args.CreateTimeEnd)
	}

	session.OrderBy("create_time desc").Limit(args.Limit, args.Offset)

	var auditList []model.Audit
	count, err := session.FindAndCount(&auditList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), auditList
}
