package common

const (
	SUCCESS = 0
)

/**************************************
 * error code
 **************************************/

const (
	// 系统域
	ErrSys        = 10001
	ErrSysReqData = 10002
	ErrSysDbExec  = 10003

	// 校验域
	ErrValidReq = 10101

	// 用户域
	ErrUserRegister = 20001
	ErrUserLogin    = 20002
	ErrUserUpdate   = 20003
	ErrUserNotExist = 20004

	ErrShopNotExist   = 20101
	ErrShopCountLimit = 20102

	ErrTaskInsert       = 20201
	ErrTaskGoodsInsert  = 20202
	ErrTaskDetailInsert = 20203

	// 消息域
	ErrSmsSend = 21001
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
	case ErrUserNotExist:
		return "商家不存在"

	case ErrShopNotExist:
		return "店铺不存在"
	case ErrShopCountLimit:
		return "店铺数量已达上限"

	case ErrTaskInsert:
		return "任务新增失败"
	case ErrTaskGoodsInsert:
		return "任务商品新增失败"
	case ErrTaskDetailInsert:
		return "任务明细新增失败"

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

func (err RespErr) Error() string {
	return err.Msg
}

func NewRespErr(code int, msg string) RespErr {
	if msg == "" {
		msg = GetMsg(code)
	}
	return RespErr{
		Code: code,
		Msg:  msg,
	}
}

/**************************************
 * 其它全局 panic 错误类型
 **************************************/

// 校验错误
type ValidErr struct {
	Msg string
}

func (err ValidErr) Error() string {
	return err.Msg
}

func NewValidErr(err error) (validErr ValidErr) {
	validErr = ValidErr{Msg: err.Error()}
	return
}

// 数据库错误
type DbErr struct {
	Msg string
}

func (err DbErr) Error() string {
	return err.Msg
}

func NewDbErr(err error) (dbErr DbErr) {
	dbErr = DbErr{Msg: err.Error()}
	return
}
