package leetcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 虽然正确实现了功能,但是时间复杂度太高, 达到O(n^3)
func p15V1(nums []int) [][]int {
	var result [][]int
	length := len(nums)
	resultMap := map[string]byte{}

	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tmpSlice := []int{nums[i], nums[j], nums[k]}
					// 使用 sort 排序防止位置不同而造成的重复
					sort.Ints(tmpSlice)
					// 使用 MAP 去重
					resultMap[fmt.Sprintf("%v,%v,%v", tmpSlice[0], tmpSlice[1], tmpSlice[2])] = 0
				}
			}
		}
	}
	for item := range resultMap {
		var tmpInt = make([]int, 3)
		// fmt.Println(item)
		for i, str := range strings.Split(item, ",") {
			// fmt.Println("index of range: ", i, str)
			j, _ := strconv.Atoi(str)
			tmpInt[i] = j
		}
		result = append(result, tmpInt)
	}
	return result

}

func p15V2(nums []int) [][]int {

	var result [][]int
	sort.Ints(nums)
	length := len(nums)
	resultMap := map[string]byte{}

	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			curSum := nums[i] + nums[j]
			if -curSum < nums[j] {
				// nums 从小到大排序， -(nums[i]+nums[j]) <= nums[j], 说明nums[j]右边已没有数能等 -（nums[i] + nums[j]）
				continue
			}

			for k := j + 1; k < length; k++ {
				if nums[k] > -curSum {
					continue
				}
				if nums[k] == -curSum {
					tmpSlice := []int{nums[i], nums[j], nums[k]}
					// 使用 MAP 去重
					resultMap[fmt.Sprintf("%v,%v,%v", tmpSlice[0], tmpSlice[1], tmpSlice[2])] = 0
				}
			}
		}
	}
	for item := range resultMap {
		var tmpInt = make([]int, 3)
		// fmt.Println(item)
		for i, str := range strings.Split(item, ",") {
			// fmt.Println("index of range: ", i, str)
			j, _ := strconv.Atoi(str)
			tmpInt[i] = j
		}
		result = append(result, tmpInt)
	}
	return result

}

// 双指针法
func p15V3(nums []int) [][]int {
	sort.Ints(nums)
	numLen := len(nums)
	resultMap := map[string]byte{}
	for i := 0; i < numLen-2; i++ {
		left, right := i+1, numLen-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum == 0 {
				// 满足条件, 使用map去重
				resultMap[fmt.Sprintf("%v,%v,%v", nums[i], nums[left], nums[right])] = 0
				left++
			} else {
				right--
			}
		}

	}

	var result [][]int

	for item := range resultMap {
		var tmpInt = make([]int, 3)
		// fmt.Println(item)
		for i, str := range strings.Split(item, ",") {
			// fmt.Println("index of range: ", i, str)
			j, _ := strconv.Atoi(str)
			tmpInt[i] = j
		}
		result = append(result, tmpInt)
	}
	return result
}
