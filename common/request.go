package common

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/gookit/validate/locales/zhcn"
)

// 验证param参数提交
func ValidateParam(g *gin.Context, format validate.MS, rule validate.MS, obj interface{}) MapItf {
	data := MapItf{}
	for k, _ := range format {
		if query, ok := g.Params.Get(k); ok && query != "" {
			query = strings.TrimSpace(query)
			data[k] = query
		}
	}

	return ValidateData(data, format, rule, &obj)
}

// 验证query参数提交
func ValidateQuery(g *gin.Context, format validate.MS, rule validate.MS, obj interface{}) MapItf {
	data := MapItf{}
	for k, _ := range format {
		if query, ok := g.GetQuery(k); ok && query != "" {
			query = strings.TrimSpace(query)
			data[k] = query
		}
	}

	return ValidateData(data, format, rule, &obj)
}

// 验证post json数据
func ValidatePostJson(g *gin.Context, format validate.MS, rule validate.MS, obj interface{}) (data MapItf) {
	jsonByte, err := g.GetRawData()
	if err != nil {
		panic(NewValidErr(err))
	}
	err = json.Unmarshal(jsonByte, &data)
	if err != nil {
		panic(NewValidErr(err))
	}
	return ValidateData(data, format, rule, &obj)
}

func ValidateData(data MapItf, format validate.MS, rule validate.MS, obj interface{}) MapItf {
	defer func() {
		validErr := recover()
		switch validErr.(type) {
		case string:
			panic(NewValidErr(errors.New(validErr.(string))))
		case error:
			panic(NewValidErr(validErr.(error)))
		}
	}()

	// 参数校验
	zhcn.RegisterGlobal()
	va := validate.Map(data)
	va.FilterRules(format)
	va.StringRules(rule)

	if !va.Validate() {
		panic(NewValidErr(errors.New(va.Errors.One())))
	}

	// 参数赋值到结构体
	err := va.BindSafeData(&obj)
	if err != nil {
		panic(NewValidErr(err))
	}
	return data
}

func (param *MapItf) FormatMapItf(rule MapItf) {
	p := *param
	for k, def := range rule {
		k = strings.TrimSpace(k)
		if p[k] != nil {
			switch def.(type) {
			case string:
				p[k] = toStr(p[k], def.(string))
			case int:
				p[k] = toInt(p[k], def.(int))
			case float64:
				p[k] = toFloat64(p[k], def.(float64))
			}
		}
	}
}

func toStr(v interface{}, def string) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return IntToStr(v.(int))
	case float64:
		return Float64ToString(v.(float64))
	}
	return def
}

func toInt(v interface{}, def int) int {
	switch v.(type) {
	case string:
		return StrToInt(v.(string), def)
	case int:
		return v.(int)
	case float64:
		return Float64ToInt(v.(float64))
	}
	return def
}

func toFloat64(v interface{}, def float64) float64 {
	switch v.(type) {
	case string:
		return StrToFloat64(v.(string), def)
	case int:
		return float64(v.(int))
	case float64:
		return v.(float64)
	}
	return def
}
