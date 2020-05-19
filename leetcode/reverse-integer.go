package leetcode

import (
	"math"
	"strconv"
)

func p7v1(x int) int {
	// 当在32平台上运行会存在bug
	restr := ""
	if x < 0 {
		x = -x
		restr = "-"
	} else if x < 10 {
		return x
	}

	s := []rune(strconv.Itoa(x))
	// 去除末尾多余的0
	for string(s[len(s)-1]) == "0" {
		s = s[0 : len(s)-1]
	}
	for i := len(s) - 1; i >= 0; i-- {
		restr += string(s[i])
	}

	// 在32位平台，进行Atoi转换时会发生截断，在之后不能判断是否发生溢出
	re, _ := strconv.Atoi(restr)

	// 题目中假设反序后的数字为有符号32位数字，发生益处时返回0
	if IsInt32Overflow(re) {
		return 0
	}
	return re
}

func p7v2(x int) int {
	// 当int是32位时也正常
	result := 0
	for x != 0 {
		i := x % 10
		x /= 10
		result *= 10
		// 判断是否溢出
		if i >= 0 && math.MaxInt32-result < i {
			return 0
		}
		if i < 0 && result-MinInt32 < i {
			return 0
		}

		result += i
	}
	return result
}
