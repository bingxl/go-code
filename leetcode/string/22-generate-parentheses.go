/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.61%)
 * Likes:    3534
 * Dislikes: 0
 * Total Accepted:    808.8K
 * Total Submissions: 1M
 * Testcase Example:  '3'
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 * Example 1:
 * Input: n = 3
 * Output: ["((()))","(()())","(())()","()(())","()()()"]
 * Example 2:
 * Input: n = 1
 * Output: ["()"]
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 8
 *
 *
 */

package string

// @lc code=start

func generateParenthesis(n int) []string {

	result := []string{}

	// left 已添加的左括号数据, right 已添加的有括号数量, cur 当前的字符串
	var dfs func(left, right int, cur string)

	dfs = func(left int, right int, cur string) {
		// 当前字符串中右括号已多于左括号; 左或右括号数量已用完则直接返回
		if left < right || left > n || right > n {
			return
		}

		if left == right && left == n {
			result = append(result, cur)
			return
		}

		dfs(left+1, right, cur+"(")
		dfs(left, right+1, cur+")")

	}

	dfs(0, 0, "")
	return result
}

// @lc code=end
