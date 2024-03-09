/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.cn/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (43.90%)
 * Likes:    4376
 * Dislikes: 0
 * Total Accepted:    1.7M
 * Total Submissions: 3.9M
 * Testcase Example:  '"()"'
 *
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * Every close bracket has a corresponding open bracket of the same type.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: s = "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: s = "(]"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * s consists of parentheses only '()[]{}'.
 *
 *
 */

// @lc code=start

package main

// 循环字符串, 遇到开始类型字符则入栈, 遇到结束类型则看栈顶字符是否和当前值匹配
// 匹配则栈顶元素出栈, 不匹配则属于无效字符.
// 循环结束后看栈是否空, 不空则不是有效字符串
func isValid(s string) bool {
	start := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	stack := []string{}

	for _, v := range s {

		if _, ok := start[string(v)]; ok {
			// 是开始类型的字符 则入栈
			stack = append(stack, string(v))
		} else {
			// 是结束类型则看和栈顶元素是否匹配
			if len(stack) == 0 {
				return false
			}
			shouldEqual := start[stack[len(stack)-1]]
			if string(v) != shouldEqual {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	// 循环结束,但栈内还有数据则有字符没未在正确的位置
	return len(stack) == 0

}

// @lc code=end
