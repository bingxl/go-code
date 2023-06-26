/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode.cn/problems/3sum/description/
 *
 * algorithms
 * Medium (37.00%)
 * Likes:    5988
 * Dislikes: 0
 * Total Accepted:    1.4M
 * Total Submissions: 3.7M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
 * k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
 *
 * 你返回所有和为 0 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 * 解释：
 * nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
 * nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
 * nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
 * 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
 * 注意，输出的顺序和三元组的顺序并不重要。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,1]
 * 输出：[]
 * 解释：唯一可能的三元组和不为 0 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0,0,0]
 * 输出：[[0,0,0]]
 * 解释：唯一可能的三元组和为 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
 *
 *
 */

// @lc code=start
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 虽然正确实现了功能,但是时间复杂度太高, 达到O(n^3)
func p15ThreeSumV1(nums []int) [][]int {
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

func p15ThreeSumV2(nums []int) [][]int {

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
func p15ThreeSumV3(nums []int) [][]int {
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

// @lc code=end
