package model

import (
	. "business/common"
)

/**
"id": "int", // 请求id
"req_no": "string", // 请求编号
"user_id": "int", // 商家编号
"router": "string", // 请求路由
"method": "string", // 请求方式
"agent": "string", // 请求头agent
"param": "string", // 请求参数
"http_code": "int", // http状态码
"code": "int", // 返回code
"msg": "string", // 返回msg
"data": "string", // 返回data
"ip": "string", // 请求ip
"server": "string", // 服务器ip
"cost": "float64", // 耗时
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type ReqLog struct {
	Id         int     `db:"id" json:"id"`
	ReqNo      string  `db:"req_no" json:"req_no"`
	UserId     int     `db:"user_id" json:"user_id"`
	Router     string  `db:"router" json:"router"`
	Method     string  `db:"method" json:"method"`
	Agent      string  `db:"agent" json:"agent"`
	Param      string  `db:"param" json:"param"`
	HttpCode   int     `db:"http_code" json:"http_code"`
	Code       int     `db:"code" json:"code"`
	Msg        string  `db:"msg" json:"msg"`
	Data       string  `db:"data" json:"data"`
	Ip         string  `db:"ip" json:"ip"`
	Server     string  `db:"server" json:"server"`
	Cost       float64 `db:"cost" json:"cost"`
	CreateTime string  `db:"create_time" json:"create_time" xorm:"created"`
	UpdateTime string  `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewReqLogModel() *ReqLog {
	return &ReqLog{}
}

func (m *ReqLog) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *ReqLog) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *ReqLog) Update(arg *ReqLog) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *ReqLog) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *ReqLog) SetId(arg int) *ReqLog {
	m.Id = arg
	return m
}

func (m *ReqLog) SetReqNo(arg string) *ReqLog {
	m.ReqNo = arg
	return m
}

func (m *ReqLog) SetUserId(arg int) *ReqLog {
	m.UserId = arg
	return m
}

func (m *ReqLog) SetRouter(arg string) *ReqLog {
	m.Router = arg
	return m
}

func (m *ReqLog) SetMethod(arg string) *ReqLog {
	m.Method = arg
	return m
}

func (m *ReqLog) SetAgent(arg string) *ReqLog {
	m.Agent = arg
	return m
}

func (m *ReqLog) SetParam(arg string) *ReqLog {
	m.Param = arg
	return m
}

func (m *ReqLog) SetHttpCode(arg int) *ReqLog {
	m.HttpCode = arg
	return m
}

func (m *ReqLog) SetCode(arg int) *ReqLog {
	m.Code = arg
	return m
}

func (m *ReqLog) SetMsg(arg string) *ReqLog {
	m.Msg = arg
	return m
}

func (m *ReqLog) SetData(arg string) *ReqLog {
	m.Data = arg
	return m
}

func (m *ReqLog) SetIp(arg string) *ReqLog {
	m.Ip = arg
	return m
}

func (m *ReqLog) SetServer(arg string) *ReqLog {
	m.Server = arg
	return m
}

func (m *ReqLog) SetCost(arg float64) *ReqLog {
	m.Cost = arg
	return m
}

func (m *ReqLog) SetCreateTime(arg string) *ReqLog {
	m.CreateTime = arg
	return m
}

func (m *ReqLog) SetUpdateTime(arg string) *ReqLog {
	m.UpdateTime = arg
	return m
}

func (m ReqLog) AsMapItf() MapItf {
	return MapItf{
		"id":          m.Id,
		"req_no":      m.ReqNo,
		"user_id":     m.UserId,
		"router":      m.Router,
		"method":      m.Method,
		"agent":       m.Agent,
		"param":       m.Param,
		"http_code":   m.HttpCode,
		"code":        m.Code,
		"msg":         m.Msg,
		"data":        m.Data,
		"ip":          m.Ip,
		"server":      m.Server,
		"cost":        m.Cost,
		"create_time": m.CreateTime,
		"update_time": m.UpdateTime,
	}
}
func (m ReqLog) Translates() map[string]string {
	return map[string]string{
		"id":          "请求id",
		"req_no":      "请求编号",
		"user_id":     "商家编号",
		"router":      "请求路由",
		"method":      "请求方式",
		"agent":       "请求头agent",
		"param":       "请求参数",
		"http_code":   "http状态码",
		"code":        "返回code",
		"msg":         "返回msg",
		"data":        "返回data",
		"ip":          "请求ip",
		"server":      "服务器ip",
		"cost":        "耗时",
		"create_time": "添加时间",
		"update_time": "更新时间",
	}
}
