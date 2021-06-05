package dao

import (
	. "business/common"
	"business/dao/model"
	"business/service/cache"
	"errors"
	"github.com/go-xorm/xorm"
)

var (
	AccountTypeMain   = "main"
	AccountTypeEmploy = "employ"
)
var AccountTypeMap = MapStr{
	AccountTypeMain:   "主账户",
	AccountTypeEmploy: "佣金账户",
}

var AccountTypeSlice = []string{AccountTypeMain, AccountTypeEmploy}

func InsertAccount(account *model.Account) {
	row := account.Insert()
	if row == 0 {
		panic(NewRespErr(ErrInsert, ""))
	}
}

func InfoAccountByUserAndType(userId int, accountType string) *model.Account {
	account := model.NewAccountModel().SetUserId(userId).SetType(accountType)
	if !account.Info() {
		InsertAccount(account)
	}
	if !account.Info() {
		panic(NewRespErr(ErrInsert, "初始化用户账户失败"))
	}
	return account
}

type UpdateAccountAmountArgs struct {
	UserId             int
	Type               string
	ChangeType         string
	AmountChange       float64
	FrozenAmountChange float64
	TaskId             string
	ShopId             int
	OrderId            int
	InOutId            string
	Remark             string
}

func UpdateAccountAmount(args UpdateAccountAmountArgs) error {
	session := DbEngine.NewSession()
	defer session.Close()
	err := UpdateAccountAmountInSess(session, args)
	if err == nil {
		if errS := session.Commit(); errS != nil {
			panic(NewDbErr(errS))
		}
	}
	return err
}

func UpdateAccountAmountInSess(s *xorm.Session, args UpdateAccountAmountArgs) error {
	account := model.NewAccountModel().SetUserId(args.UserId).SetType(args.Type)
	if !account.Info() {
		return errors.New("账户信息有误")
	}

	set := model.NewAccountModel().
		SetAmount(account.Amount + args.AmountChange).
		SetFrozenAmount(account.FrozenAmount + args.FrozenAmountChange)

	row, err := s.Cols("amount", "frozen_amount").Update(set, account)
	if err != nil {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewDbErr(err))
	}
	if row == 0 {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		return errors.New("更新账户金额失败")
	}

	log := &model.AccountLog{
		AccountId: account.Id,
		UserId:    account.UserId,
		Type:      args.ChangeType,
		AmountOld: account.Amount,
		AmountNew: account.Amount + args.AmountChange,
		FrozenOld: account.FrozenAmount,
		FrozenNew: account.FrozenAmount + args.FrozenAmountChange,
		TaskId:    args.TaskId,
		ShopId:    args.ShopId,
		OrderId:   args.OrderId,
		InOutId:   args.InOutId,
		Remark:    args.Remark,
	}
	row, err = s.Insert(log)
	if err != nil {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewDbErr(err))
	}
	if row == 0 {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		return errors.New("新增资金记录失败")
	}

	cache.NewCacheUserInfo(account.UserId).DeleteCacheUserInfo()
	return nil
}
