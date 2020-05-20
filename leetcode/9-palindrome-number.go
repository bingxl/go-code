package leetcode

import "strconv"

// 判断数字是否是回文

func p9v1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	re := strconv.Itoa(x)
	half := len(re) / 2
	tail := len(re) - 1
	for i := 0; i < half; i++ {
		if re[i] != re[tail-i] {
			return false
		}
	}
	return true
}
