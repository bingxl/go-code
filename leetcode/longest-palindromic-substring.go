package leetcode

import (
	"math"
)

// P5v1输出最长回文串
func p5v1(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	ru := []rune(s)
	n := len(ru)
	imax, imin := 0, 0
	for i := 1; i < n; i++ {
		curMax, curMin := i, i

		// 处理重复连续重复的字符串
		for j := curMin - 1; j >= 0; j-- {
			if j >= 0 && ru[j] == ru[i] {
				curMin = j
			} else {
				break
			}
		}
		for k := curMax + 1; k < n; k++ {
			if ru[k] == ru[i] {
				curMax = k
			} else {
				break
			}
		}

		// 跳过连续字符串
		i = int(math.Max(float64(curMax-1), float64(i)))

		for j, k := curMin-1, curMax+1; j >= 0 && k < n; {
			if ru[j] == ru[k] {
				curMin, curMax = j, k
				j--
				k++
			} else {
				break
			}
		}
		if curMax-curMin > imax-imin {
			imax, imin = curMax, curMin
		}
	}

	return s[imin : imax+1]
}
