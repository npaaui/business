package common

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

/****************************
 * 时间相关
 ****************************/
func GetNow() (date string) {
	date = time.Now().Format("2006-01-02 15:04:05")
	return
}

func GetForever() (date string) {
	date = "2099-01-01 00:00:00"
	return
}

func GetTomorrowBegin() (date string) {
	date = time.Now().AddDate(0, 0, 1).Format("2006-01-02 00:00:00")
	return
}

func GetAfterHour(hour int) (date string) {
	date = time.Now().Add(time.Hour * time.Duration(hour)).Format("2006-01-02 15:04:05")
	return
}

/****************************
 * 加密相关
 ****************************/
// 获取hash
func GetHash(str string) (hash string) {
	h := md5.Sum([]byte(str))
	hash = fmt.Sprintf("%x", h)
	return
}

// 获取随机数
func RandNumString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		rs = append(rs, strconv.Itoa(rand.Intn(10)))
	}
	return strings.Join(rs, "")
}

/****************************
 * 数据格式转换
 ****************************/
func StrToInt(str string, def int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		Log(LogLevelWarning, fmt.Errorf("字符串「"+str+"」转整数失败: ", err))
		return def
	}
	return i
}

func Float64ToInt(f float64) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", f))
	return i
}

func IntToStr(i int) string {
	str := strconv.Itoa(i)
	return str
}

func Float64ToString(f float64) string {
	str := strconv.FormatFloat(f, 'g', -1, 64)
	decimalNum, err := decimal.NewFromString(str)
	if err != nil {
		return str
	}
	str = decimalNum.String()
	return str
}

func StrToFloat64(s string, def float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	fmt.Printf("失败:%f", f)
	if err != nil {
		Log(LogLevelWarning, fmt.Errorf("字符串「"+s+"」转浮点数失败: ", err))
		return def
	}
	return f
}
