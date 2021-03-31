package common

type Arr map[string]interface{}

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
