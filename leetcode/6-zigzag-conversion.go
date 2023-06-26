/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] N 字形变换
 *
 * https://leetcode.cn/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (52.04%)
 * Likes:    2069
 * Dislikes: 0
 * Total Accepted:    567.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
 *
 * 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
 *
 * 请你实现这个将字符串进行指定行数变换的函数：
 *
 *
 * string convert(string s, int numRows);
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "PAYPALISHIRING", numRows = 3
 * 输出："PAHNAPLSIIGYIR"
 *
 * 示例 2：
 *
 *
 * 输入：s = "PAYPALISHIRING", numRows = 4
 * 输出："PINALSIGYAHRPI"
 * 解释：
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "A", numRows = 1
 * 输出："A"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 由英文字母（小写和大写）、',' 和 '.' 组成
 * 1
 *
 *
 */

// @lc code=start
package main

// 思路：使用numRows个切片来存储字符，注意正向或反向存储的问题

// 0     6      12
// 1   5 7   11 13
// 2 4   8 10   14
// 3     9      15
func convert(s string, numRows int) string {
	// 只需要一行时直接返回
	if numRows <= 1 {
		return s
	}
	ru := []rune(s)
	re := make([]([]rune), numRows)
	reverse := false
	for k, v := range ru {
		i := k % (numRows - 1)
		j := i
		// 第一列最后一行的值开始反向存储
		if k != 0 && i == 0 {
			reverse = !reverse
		}

		if reverse {
			j = numRows - 1 - i
		}

		re[j] = append(re[j], v)
	}

	var result string
	for _, sli := range re {
		for _, v := range sli {
			result += string(v)
		}
	}
	// fmt.Println(ru)
	// fmt.Println(re)
	return result
}

// @lc code=end
