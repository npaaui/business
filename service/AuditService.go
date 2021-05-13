package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
	"reflect"
)

type AuditService struct{}

func NewAuditService() *AuditService {
	return &AuditService{}
}

// 审核后续执行方法
var AuditAfterFuncArr = MapItf{
	"recharge": AuditAfterRecharge,
	"withdraw": AuditAfterWithdraw,
}

/**
 * 审核列表
 */
func (s *AuditService) ListAudit(args *dao.ListAuditArgs) (data *RespList) {
	count, list := dao.ListAudit(args)
	data = NewRespList(count, list)
	return
}

/**
 * 更新审核
 */
func (s *AuditService) UpdateAudit(args *model.Audit) {
	audit := model.NewAuditModel().SetId(args.Id)
	if !audit.Info() {
		panic(NewRespErr(ErrNotExist, "审核信息有误"))
	}

	// 获取审核动作
	auditAction := model.NewAuditActionModel().SetCode(audit.Action)
	if !auditAction.Info() {
		panic(NewRespErr(ErrNotExist, "无效审核动作"))
	}

	// 执行审核后续动作
	function := AuditAfterFuncArr[audit.Action]
	f := reflect.ValueOf(function)
	in := []reflect.Value{reflect.ValueOf(args)}
	f.Call(in)

	if row := audit.Update(args); row == 0 {
		panic(NewRespErr(ErrUpdate, ""))
	}
}

func AuditAfterRecharge(m *model.Audit) bool {
	if m.Status == dao.AuditStatusPass {
		// 获取充值金额
		accountInOut := model.NewAccountInOutModel().SetId(m.LinkId)
		if !accountInOut.Info() {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n 审核异常原因：充值记录未找到。"
		}

		// 充值
		err := dao.UpdateAccountAmount(dao.UpdateAccountAmountArgs{
			UserId:       m.UserId,
			Type:         dao.AccountTypeMain,
			ChangeType:   dao.AccountInOutTypeRecharge,
			AmountChange: accountInOut.Amount,
			InOutId:      m.LinkId,
		})
		if err != nil {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n " + err.Error()
		}
		m.Status = dao.AuditStatusPass
	} else {
		m.Status = dao.AuditStatusFail
	}
	return true
}

func AuditAfterWithdraw(m *model.Audit) bool {
	if m.Status == dao.AuditStatusPass {
		// 获取提现金额
		accountInOut := model.NewAccountInOutModel().SetId(m.LinkId)
		if !accountInOut.Info() {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n 审核异常原因：充值记录未找到。"
		}

		// 获取账户金额
		account := dao.InfoAccountByUserAndType(accountInOut.UserId, dao.AccountTypeMain)
		if account.Amount < accountInOut.Amount {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n 审核异常原因：提现金额高于账户余额。"
		}

		// 提现
		err := dao.UpdateAccountAmount(dao.UpdateAccountAmountArgs{
			UserId:       m.UserId,
			Type:         dao.AccountTypeMain,
			ChangeType:   dao.AccountInOutTypeWithdraw,
			AmountChange: -accountInOut.Amount,
			InOutId:      m.LinkId,
		})
		if err != nil {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n " + err.Error()
		}
		m.Status = dao.AuditStatusPass
	} else {
		m.Status = dao.AuditStatusFail
	}
	return true
}
