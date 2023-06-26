/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode.cn/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (60.28%)
 * Likes:    4345
 * Dislikes: 0
 * Total Accepted:    994.2K
 * Total Submissions: 1.6M
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
 *
 * 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 返回容器可以储存的最大水量。
 *
 * 说明：你不能倾斜容器。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 * 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 * 示例 2：
 *
 *
 * 输入：height = [1,1]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 2 <= n <= 10^5
 * 0 <= height[i] <= 10^4
 *
 *
 */

// @lc code=start

package main

import (
	"fmt"
	"math"
)

// 暴力破解, 两次循环, O(n^2)
func p11MaxAreaV1(height []int) int {
	maxArea := 0.0

	for i := 0; (i + 1) < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := math.Min(float64(height[i]), float64(height[j])) * float64(j-i)

			if maxArea < area {
				maxArea = area
			}
		}
	}
	return int(maxArea)
}

// 从两侧搜寻
func p11MaxAreaV2(height []int) int {
	maxArea := 0

	i, j := 0, len(height)-1

	for i != j {
		min := i
		if height[i] > height[j] {
			// j位置元素小则j左移
			min = j
			j--
		} else {
			// i位置元素小,则i右移
			i++
		}
		// j-i+1是因为i/j的自增自减在前面
		area := height[min] * (j - i + 1)
		if maxArea < area {
			maxArea = area
		}
		fmt.Printf("[%v, %v], max: %v, curMax: %v \n", i, j, maxArea, area)

	}
	fmt.Printf("################end one case###########")

	return maxArea
}

// @lc code=end
