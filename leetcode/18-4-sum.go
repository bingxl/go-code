/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.cn/problems/4sum/description/
 *
 * algorithms
 * Medium (36.66%)
 * Likes:    1877
 * Dislikes: 0
 * Total Accepted:    559.9K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers, return an array of all the unique
 * quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
 *
 *
 * 0 <= a, b, c, d < n
 * a, b, c, and d are distinct.
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 *
 *
 * You may return the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,0,-1,0,-2,2], target = 0
 * Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,2,2,2,2], target = 8
 * Output: [[2,2,2,2]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 *
 *
 */

// @lc code=start
package main

import (
	"fmt"
	"sort"
)

// 四重循环查找
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)

	resMap := make(map[string]bool)
	numsLen := len(nums)
	res := [][]int{}
	for a := 0; a < numsLen-3; a++ {
		for b := a + 1; b < numsLen-2; b++ {
			for c := b + 1; c < numsLen-1; c++ {
				for d := c + 1; d < numsLen; d++ {
					sum := nums[a] + nums[b] + nums[c] + nums[d]
					// sums 按 asc 排序过, sum > target 则后面不会有满足条件的值
					if sum > target {
						break
					}
					if sum == target {
						key := fmt.Sprintf("%v,%v,%v,%v", nums[a], nums[b], nums[c], nums[d])
						if _, exists := resMap[key]; !exists {
							resMap[key] = true
							res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
						}
						// 已找到一组满足条件的, 则 d 没必要继续查找,就算有满足条件的值也只会是重复当前的值
						break
					}
				}
			}
		}
	}

	return res

}

// 三重循环 + map
func fourSum1(nums []int, target int) [][]int {
	numsLen := len(nums)
	if numsLen < 4 {
		return [][]int{}
	}

	sort.Ints(nums)
	res := [][]int{}
	resMap := make(map[string]bool)

	numsMap := make(map[int]int, numsLen)
	for i, v := range nums {
		numsMap[v] = i
	}

	for a := 0; a < numsLen-3; a++ {
		for b := a + 1; b < numsLen-2; b++ {
			for c := b + 1; c < numsLen-1; c++ {
				sum := nums[a] + nums[b] + nums[c]
				if sum > target {
					break
				}
				if v, ok := numsMap[target-sum]; ok && v > c {
					key := fmt.Sprintf("%v-%v-%v-%v", nums[a], nums[b], nums[c], nums[v])
					if _, ok := resMap[key]; !ok {
						resMap[key] = true
						res = append(res, []int{nums[a], nums[b], nums[c], nums[v]})
					}
				}

			}
		}
	}
	return res
}

// @lc code=end
