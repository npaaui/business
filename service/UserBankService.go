package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
)

func (s *UserService) ListUserBank(args *dao.ListUserBankArgs) (data *RespList) {
	count, list := dao.ListUserBank(args)
	data = NewRespList(count, list)
	return
}

func (s *UserService) InsertUserBank(userBank *model.UserBank) {
	// 校验银行
	category := model.NewCategoryModel().SetId(userBank.BankCategoryId)
	if !NewConfigService().InfoCategory(category) {
		panic(NewRespErr(ErrNotExist, "不存在银行"))
	}
	userBank.BankCategoryName = category.Name

	userBankS := &model.UserBank{
		Code: userBank.Code,
	}
	_ = userBankS.Info()
	if userBankS.Id > 0 {
		panic(NewRespErr(ErrInsert, "该银行卡记录已存在"))
	}
	ret := userBank.SetCreateTime(GetNow()).SetUpdateTime(GetNow()).Insert()
	if ret != 1 {
		panic(NewRespErr(ErrInsert, ""))
	}
}

func (s *UserService) UpdateUserBank(set *model.UserBank) {
	userBank := &model.UserBank{
		Id:     set.Id,
		UserId: set.UserId,
	}
	if !userBank.Info() {
		panic(NewRespErr(ErrNotExist, "不存在的银行卡记录"))
	}

	// 校验银行
	category := model.NewCategoryModel().SetId(set.BankCategoryId)
	if !NewConfigService().InfoCategory(category) {
		panic(NewRespErr(ErrNotExist, "不存在银行"))
	}
	set.BankCategoryName = category.Name

	row := userBank.Update(set)
	if row == 0 {
		panic(NewRespErr(ErrUpdate, ""))
	}
}

func (s *UserService) DeleteUserBank(set *model.UserBank) {
	userBank := &model.UserBank{
		Id:     set.Id,
		UserId: set.UserId,
	}
	if !userBank.Info() {
		panic(NewRespErr(ErrNotExist, "不存在的银行卡记录"))
	}

	row := userBank.Delete()
	if row == 0 {
		panic(NewRespErr(ErrDelete, ""))
	}
}
