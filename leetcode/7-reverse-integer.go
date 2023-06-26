/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode.cn/problems/reverse-integer/description/
 *
 * algorithms
 * Medium (35.37%)
 * Likes:    3841
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 3.4M
 * Testcase Example:  '123'
 *
 * 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
 *
 * 如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
 * 假设环境不允许存储 64 位整数（有符号或无符号）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：x = 123
 * 输出：321
 *
 *
 * 示例 2：
 *
 *
 * 输入：x = -123
 * 输出：-321
 *
 *
 * 示例 3：
 *
 *
 * 输入：x = 120
 * 输出：21
 *
 *
 * 示例 4：
 *
 *
 * 输入：x = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * -2^31
 *
 *
 */

// @lc code=start
package main

import (
	"math"
	"strconv"
)

func p7Reverse(x int) int {
	// 当在32平台上运行会存在bug
	restr := ""
	if x < 0 {
		x = -x
		restr = "-"
	} else if x < 10 {
		return x
	}
	// IsInt32Overflow 判断 x是否溢出有符号32位
	IsInt32Overflow := func(x int) bool {
		if x < math.MaxInt32 && x > -math.MaxInt32-1 {
			return false
		}
		return true
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

func p7ReverseV2(x int) int {
	// 当int是32位时也正常
	result := 0
	// MinInt32 最小有符号32位
	var MinInt32 = -math.MaxInt32 - 1

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

// @lc code=end
