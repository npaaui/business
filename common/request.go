package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/gookit/validate/locales/zhcn"
)

// 验证param参数提交
func ValidateParam(g *gin.Context, format validate.MS, rule validate.MS, obj interface{}) MapItf {
	data := MapItf{}
	for k := range format {
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
	for k := range format {
		if query, ok := g.GetQuery(k); ok && query != "" {
			query = strings.TrimSpace(query)
			data[k] = query
		}
	}

	return ValidateData(data, format, rule, &obj)
}

// 验证 post form 参数提交
func ValidatePostForm(g *gin.Context, format validate.MS, rule validate.MS, obj interface{}) MapItf {
	data := MapItf{}
	for k := range format {
		if query, ok := g.GetPostForm(k); ok && query != "" {
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
	// int, float类型值传了空字符串过不了filter，暂时做一下处理
	for k, v := range format {
		if v == "int" && data[k] == "" {
			data[k] = 0
		}
		if v == "float" && data[k] == "" {
			data[k] = "0"
		}
	}

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

func LoadPostFile(g *gin.Context, fileKey string, resType string) string {
	header, err := g.FormFile(fileKey)
	if err != nil {
		panic(NewValidErr(err))
	}
	fileName := header.Filename

	fileExt := path.Ext(fileName)
	fileTime := time.Now().Format("20060102130405")
	fileRand := IntToStr(rand.Intn(100))
	name := fileTime + fileRand + "[" + strings.TrimRight(fileName, "."+fileExt) + "]" + fileExt

	postFile, _ := header.Open()
	defer func() {
		err = postFile.Close()
		if err != nil {
			panic(NewSysErr(fmt.Errorf("postFile文件句柄关闭失败:%w", err)))
		}
	}()

	src := "upload/" + resType
	upload, err := os.Create(src + "/" + name)
	defer func() {
		err = upload.Close()
		if err != nil {
			panic(NewSysErr(fmt.Errorf("upload文件句柄关闭失败:%w", err)))
		}
	}()

	if os.IsNotExist(err) {
		// 判断upload文件夹是否存在
		_, err = os.Stat(src)
		if os.IsNotExist(err) {
			err = os.MkdirAll(src, os.ModePerm)
			if err != nil {
				panic(NewSysErr(fmt.Errorf("upload文件夹创建失败:%w", err)))
			}
			upload, _ = os.Create(src + "/" + name)
		} else {
			panic(NewSysErr(fmt.Errorf(name+"文件创建失败:%w", err)))
		}
	}

	_, err = io.Copy(upload, postFile)
	if err != nil {
		panic(NewSysErr(fmt.Errorf(name+"文件创建失败:%w", err)))
	}
	return resType + "/" + name
}
