package dao

import (
	. "business/common"
	"business/dao/model"
	"errors"
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

func InsertAudit(audit *model.Audit) error {
	session := DbEngine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}

	row, err := session.Insert(audit)
	if err != nil {
		_ = session.Rollback()
		return err
	}
	if row == 0 {
		_ = session.Rollback()
		return errors.New("审核记录未新增")
	}

	has, err := session.Get(audit)
	if err != nil {
		_ = session.Rollback()
		return err
	}
	if !has {
		_ = session.Rollback()
		return errors.New("审核记录未新增")
	}

	// 添加审核日志
	row, err = session.Insert(&model.AuditLog{
		AuditId: audit.Id,
		UserId:  audit.UserId,
		Status:  audit.Status,
		LinkId:  audit.LinkId,
		Remark:  audit.Remark,
		OpsId:   audit.OpsId,
	})
	if err != nil {
		_ = session.Rollback()
		return err
	}
	if row == 0 {
		_ = session.Rollback()
		return errors.New("审核日志未新增")
	}
	_ = session.Commit()
	return nil
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

	session.OrderBy("a.create_time desc").Limit(args.Limit, args.Offset)

	var auditList []model.Audit
	count, err := session.FindAndCount(&auditList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), auditList
}
