package dao

import (
	"strings"

	. "business/common"
)

/**
 * AddSlashes() 函数返回在预定义字符之前添加反斜杠的字符串。
 * 预定义字符是：
 * 单引号（'）
 * 双引号（"）
 * 反斜杠（\）
 */
func AddSlashes(str string) string {
	var tmpRune []rune
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

/*
 * StripSlashes() 函数删除由 AddSlashes() 函数添加的反斜杠。
 */
func StripSlashes(str string) string {
	var dstRune []rune
	strRune := []rune(str)
	strLength := len(strRune)
	for i := 0; i < strLength; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}

/**
 * where in多值查询 string
 */
func WhereInString(arg []string) string {
	if len(arg) == 0 {
		return ""
	}
	safeStr := "("
	for _, v := range arg {
		safeStr += "'" + AddSlashes(v) + "',"
	}
	return strings.TrimRight(safeStr, ",") + ")"
}

/**
 * where in多值查询 int
 */
func WhereInInt(arg []int) string {
	if len(arg) == 0 {
		return ""
	}
	safeStr := "("
	for _, v := range arg {
		safeStr += AddSlashes(IntToStr(v)) + ","
	}
	return strings.TrimRight(safeStr, ",") + ")"
}
