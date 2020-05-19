package leetcode

import "math"

// MinInt32 最小有符号32位
var MinInt32 = -math.MaxInt32 - 1

// IsInt32Overflow 判断 x是否溢出有符号32位
func IsInt32Overflow(x int) bool {
	if x < math.MaxInt32 && x > -math.MaxInt32-1 {
		return false
	}
	return true
}
