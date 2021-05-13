package service

import (
	"errors"

	. "business/common"
	"business/dao/model"
)

type SmsService struct{}

func NewSmsService() *SmsService {
	return &SmsService{}
}

var SmsValidTypeRegister = "register"                               // 注册
var SmsValidTypeUpdatePassword = "update_password"                  //修改密码
var SmsValidTypeUpdateWithdrawPassword = "update_withdraw_password" //修改提现密码
var SmsValidTypeArr = []string{SmsValidTypeRegister, SmsValidTypeUpdatePassword, SmsValidTypeUpdateWithdrawPassword}

func SendSmsValid(smsValid *model.SmsValid) error {
	row := smsValid.SetCode(RandNumString(4)).
		SetCreateTime(GetNow()).
		SetUpdateTime(GetNow()).
		SetExpireTime(GetNow()).Insert()
	if row == 0 {
		return errors.New("生成短信验证码失败")
	}
	return nil
}
