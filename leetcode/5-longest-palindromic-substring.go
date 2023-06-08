/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode.cn/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (37.58%)
 * Likes:    6556
 * Dislikes: 0
 * Total Accepted:    1.4M
 * Total Submissions: 3.8M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 * 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由数字和英文字母组成
 *
 *
 */

// @lc code=start
package leetcode

import (
	"math"
)

// P5v1输出最长回文串
func longestPalindrome(s string) string {
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

// @lc code=end
