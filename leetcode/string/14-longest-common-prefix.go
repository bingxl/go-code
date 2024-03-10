/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode.cn/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (43.34%)
 * Likes:    2804
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 2.5M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：strs = ["flower","flow","flight"]
 * 输出："fl"
 *
 *
 * 示例 2：
 *
 *
 * 输入：strs = ["dog","racecar","car"]
 * 输出：""
 * 解释：输入不存在公共前缀。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] 仅由小写英文字母组成
 *
 *
 */

// @lc code=start
package string

import (
	"strings"
)

func longestCommonPrefixV1(strs []string) string {
	if (len(strs)) == 0 {
		return ""
	}
	minLen := len(strs[0])
	commonPrefix := strings.Builder{}

	for _, s := range strs {
		if minLen > len(s) {
			minLen = len(s)
		}
	}

	for i := 0; i < minLen; i++ {
		m := make(map[byte]int)
		var commonByte byte
		for _, s := range strs {
			commonByte = s[i]
			m[commonByte]++
		}
		if len(m) != 1 {
			break
		} else {
			commonPrefix.WriteByte(commonByte)
		}
	}

	return commonPrefix.String()
}

func longestCommonPrefixV2(strs []string) string {
	// 如果为空,返回"", 如果只有一个字符串则返回这个字符串, 因为后面遍历比较时可以不用和strs[0]比较
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	baseStr := strs[0]

	for _, s := range strs[1:] {
		// 如果s没有对应的baseStr前缀则不断缩短baseStr
		for strings.Index(s, baseStr) != 0 {
			// 基串为空,直接返回
			if baseStr == "" {
				return baseStr
			}
			baseStr = baseStr[:len(baseStr)-1]
		}
	}
	return baseStr
}

// @lc code=end
