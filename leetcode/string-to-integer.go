package leetcode

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

// 从第一个非空字符到第一个非数字字符（除开头的负号）
func p8v1(str string) int {
	// 去除空格
	str = strings.TrimLeft(str, " ")
	tmp := []rune(str)
	if len(tmp) == 0 {
		return 0
	}
	// 如果开头不是 +/-/数字 则返回0
	if str[0] != '+' && str[0] != '-' && !unicode.IsNumber(tmp[0]) {
		return 0
	}
	i := 1
	for i < len(tmp) {
		// 到第一个非数字字符结束
		if !unicode.IsNumber(tmp[i]) {
			break
		} else {
			i++
		}
	}
	// 将字符转换为整数

	result, _ := strconv.Atoi(str[0:i])
	// 溢出32位数字部分返回32位的最大或最小值
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < MinInt32 {
		return MinInt32
	}
	return result
}
