// cSpell: disable
/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (39.11%)
 * Likes:    9221
 * Dislikes: 0
 * Total Accepted:    2.4M
 * Total Submissions: 6M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */

// @lc code=start
// cSpell: enable
package string

import (
	"math"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var subStr string
	var maxLen float64 = 0
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if strings.Contains(subStr, char) {
			maxLen = math.Max(maxLen, float64(len(subStr)))
			i -= len(subStr)
			subStr = ""
		} else {
			subStr = subStr + char
		}
	}

	return int(math.Max(maxLen, float64(len(subStr))))
}

func lengthOfLongestSubstringV1(s string) int {
	maxLength := 0.0

	isUnique := func(s string, start, end int) bool {
		store := make(map[string]int)
		for i := start; i < end; i++ {
			char := string(s[i])
			if store[char] == 1 {
				return false
			}
			store[char] = 1
		}
		return true
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isUnique(s, i, j) {
				maxLength = math.Max(maxLength, float64(j-i))
			}
		}
	}
	return int(maxLength)
}

// @TODO 此解法有问题
func lengthOfLongestSubstringV2(s string) int {
	n := len(s)
	hash := make(map[string]bool)
	i, j, ans := 0, 0, 0.0
	for i < n && j < n {
		char := string(s[j])
		// fmt.Println(i, j)
		if hash[char] {
			hash[char] = true
			j++
			ans = math.Max(ans, float64(j-i))
		} else {
			hash[string(s[i])] = false
			i++
		}
	}
	return int(ans)
}

func lengthOfLongestSubstringV3(s string) int {
	n := len(s)
	hash := make(map[string]float64)
	ans := 0.0
	for i, j := 0, 0; j < n; j++ {
		char := string(s[j])
		if v, ok := hash[char]; ok {
			i = int(math.Max(v, float64(i)))
		}
		ans = math.Max(ans, float64(j-i+1))
		hash[char] = float64(j + 1)
	}
	return int(ans)
}

// @lc code=end
