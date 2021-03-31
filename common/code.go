package common

const (
	SUCCESS = 0
)

const (
	/**
	 * 全局错误
	 */
	// 系统域
	ErrSys = 10001
	ErrSysReqData = 10002
	ErrSysDbExec = 10003

	// 校验域
	ErrValidReq = 10101

	/**
	 * 业务错误
	 */
	// 用户域
	ErrUserRegister = 20101
	ErrUserLogin = 20102
	ErrUserUpdate = 20103

	// 消息域
	ErrSmsSend = 20201
)

func GetMsg(code int) string {
	switch code {
	case SUCCESS:
		return "成功"

	case ErrSys:
		return "系统繁忙 请稍后重试"
	case ErrSysReqData:
		return "请求数据解析失败 请稍后重试"
	case ErrSysDbExec:
		return "数据错误"

	case ErrValidReq:
		return "参数错误"

	case ErrUserRegister:
		return "注册失败"
	case ErrUserLogin:
		return "登录失败"
	case ErrUserUpdate:
		return "修改密码失败"

	case ErrSmsSend:
		return "发送失败"
	default:
		return "未知错误"
	}
}

type RespErr struct {
	Code int
	Msg  string
}

func NewRespErr(code int, msg string) *RespErr {
	if msg == "" {
		msg = GetMsg(code)
	}
	return &RespErr{
		Code: code,
		Msg:  msg,
	}
}

func (ec *RespErr) SetMsg(msg string) *RespErr {
	ec.Msg = msg
	return ec
}
