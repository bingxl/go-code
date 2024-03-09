/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (59.18%)
 * Likes:    2772
 * Dislikes: 0
 * Total Accepted:    827.3K
 * Total Submissions: 1.4M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent. Return the answer in
 * any order.
 *
 * A mapping of digits to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 * Example 1:
 *
 *
 * Input: digits = "23"
 * Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * Example 2:
 *
 *
 * Input: digits = ""
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: digits = "2"
 * Output: ["a","b","c"]
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] is a digit in the range ['2', '9'].
 *
 *
 */

// @lc code=start
package main

// 回溯法
// 遍历digits, 将当前digit对应的所有字母分别月已有字符组合
// 调用次数多时可将maps, combine 函数提到letterCombinations 外

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// 数字对应的字母
	var maps = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	results := []string{}

	var combine func(index int, combineStr string)
	combine = func(index int, combineStr string) {
		if index == len(digits) {
			results = append(results, combineStr)
			return
		}
		for _, char := range maps[digits[index]] {

			combine(index+1, combineStr+string(char))
		}
	}

	combine(0, "")
	return results
}

// @lc code=end
