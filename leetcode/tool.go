package leetcode

/**
leetcode包下共用工具
*/
import (
	"fmt"
	"math"
	"strconv"
)

const checkMark = "\u2713" // 对号 %v
const ballotX = "\u2717"   // 错号

// MinInt32 最小有符号32位
var MinInt32 = -math.MaxInt32 - 1

// IsInt32Overflow 判断 x是否溢出有符号32位
func IsInt32Overflow(x int) bool {
	if x < math.MaxInt32 && x > -math.MaxInt32-1 {
		return false
	}
	return true
}

// SliceToString 将slice转为string
func SliceToString(sli []interface{}) string {
	str := ""
	for _, v := range sli {
		tmpStr := ""
		switch v.(type) {
		case int:
			tmpStr = strconv.Itoa(v.(int))
		case string:
			tmpStr = v.(string)
		}
		str += fmt.Sprintf("->%s", tmpStr)
	}
	return str
}
