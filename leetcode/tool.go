package main

/**
leetcode包下共用工具
*/
import (
	"fmt"
	"slices"
	"sort"
	"strconv"
)

const rightSymbol = "   \033[32m\u2713\033[0m" // 绿色的对号 %v
const errorSymbol = "   \033[31m\u2717\033[0m" // 红色的错号

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

// 比较两个切片是否相同，不考虑顺序
func compareSlicesWithDisorder(slice1, slice2 []int) bool {

	// 排序切片
	sort.Ints(slice1)
	sort.Ints(slice2)

	return slices.Equal(slice1, slice2)
}
