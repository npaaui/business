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
	"task":     AuditAfterTask,
}

/**
 * 审核列表
 */
func (s *AuditService) ListAudit(args *dao.ListAuditArgs) (data *RespList) {
	count, list := dao.ListAudit(args)
	var resList = make([]map[string]interface{}, len(list))
	for k, v := range list {
		item := v.AsMapItf()
		item["status_desc"] = dao.AuditStatusMap[v.Status]
		resList[k] = item
	}
	data = NewRespList(count, resList)
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

	if audit.Status != dao.AuditStatusInit {
		panic(NewRespErr(ErrNotExist, "不可重复审核"))
	}

	// 获取审核动作
	auditAction := model.NewAuditActionModel().SetCode(audit.Action)
	if !auditAction.Info() {
		panic(NewRespErr(ErrNotExist, "无效审核动作"))
	}

	// 执行审核后续动作
	setAudit := model.NewAuditModel()
	*setAudit = *audit
	setAudit.SetStatus(args.Status).SetRemark(args.Remark)
	function := AuditAfterFuncArr[audit.Action]
	f := reflect.ValueOf(function)
	in := []reflect.Value{reflect.ValueOf(setAudit)}
	ret := f.Call(in)

	if row := audit.Update(setAudit); row == 0 {
		panic(NewRespErr(ErrUpdate, ""))
	}

	if !ret[0].Bool() {
		panic(NewRespErr(ErrAuditStop, setAudit.Remark))
	}
}

func AuditAfterRecharge(m *model.Audit) bool {
	if m.Status == dao.AuditStatusPass {
		// 获取充值金额
		accountInOut := model.NewAccountInOutModel().SetId(StrToInt(m.LinkId, 0))
		if !accountInOut.Info() {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n 审核异常原因：充值记录未找到。"
			return false
		}

		// 充值
		err := dao.UpdateAccountAmount(dao.UpdateAccountAmountArgs{
			UserId:       m.UserId,
			Type:         dao.AccountTypeMain,
			ChangeType:   dao.AccountInOutTypeRecharge,
			AmountChange: accountInOut.Amount,
			InOutId:      StrToInt(m.LinkId, 0),
		})
		if err != nil {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n " + err.Error()
			return false
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
		accountInOut := model.NewAccountInOutModel().SetId(StrToInt(m.LinkId, 0))
		if !accountInOut.Info() {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n 审核异常原因：充值记录未找到。"
			return false
		}

		// 获取账户金额
		account := dao.InfoAccountByUserAndType(accountInOut.UserId, dao.AccountTypeMain)
		if account.Amount < accountInOut.Amount {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n 审核异常原因：提现金额高于账户余额。"
			return false
		}

		// 提现
		err := dao.UpdateAccountAmount(dao.UpdateAccountAmountArgs{
			UserId:       m.UserId,
			Type:         dao.AccountTypeMain,
			ChangeType:   dao.AccountInOutTypeWithdraw,
			AmountChange: -accountInOut.Amount,
			InOutId:      StrToInt(m.LinkId, 0),
		})
		if err != nil {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n 提现失败: " + err.Error()
			return false
		}
		m.Status = dao.AuditStatusPass
	} else {
		m.Status = dao.AuditStatusFail
	}
	return true
}

func AuditAfterTask(m *model.Audit) bool {
	task := model.NewTaskModel().SetId(StrToInt64(m.LinkId, 0))
	if !task.Info() {
		m.Status = dao.AuditStatusStop
		m.Remark += "\n 审核异常：任务记录未找到。"
		return false
	}

	// 获取冻结金额
	account := dao.InfoAccountByUserAndType(m.UserId, dao.AccountTypeMain)
	if account.FrozenAmount < task.PayAmount {
		m.Status = dao.AuditStatusStop
		m.Remark += "\n 审核异常：冻结金额小于任务金额。"
		return false
	}

	if m.Status == dao.AuditStatusPass {
		// 支付任务金额
		err := dao.UpdateAccountAmount(dao.UpdateAccountAmountArgs{
			UserId:             m.UserId,
			Type:               dao.AccountTypeMain,
			ChangeType:         dao.AccountInOutTypeTask,
			FrozenAmountChange: -task.PayAmount,
			TaskId:             StrToInt64(m.LinkId, 0),
			ShopId:             task.ShopId,
		})
		if err != nil {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n 支付金额失败: " + err.Error()
			return false
		}

		err = NewOrderService().InitOrders(StrToInt64(m.LinkId, 0))
		if err != nil {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n 添加订单失败: " + err.Error()
			return false
		}
		m.Status = dao.AuditStatusPass
	} else {
		// 解冻金额
		err := dao.UpdateAccountAmount(dao.UpdateAccountAmountArgs{
			UserId:             m.UserId,
			Type:               dao.AccountTypeMain,
			ChangeType:         dao.AccountInOutTypeTask,
			AmountChange:       task.PayAmount,
			FrozenAmountChange: -task.PayAmount,
			TaskId:             StrToInt64(m.LinkId, 0),
			ShopId:             task.ShopId,
		})
		if err != nil {
			m.Status = dao.AuditStatusStop
			m.Remark += "\n 解冻金额失败 " + err.Error()
			return false
		}
		m.Status = dao.AuditStatusFail
	}
	return true
}
