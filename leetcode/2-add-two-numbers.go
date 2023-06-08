/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode.cn/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (42.41%)
 * Likes:    9647
 * Dislikes: 0
 * Total Accepted:    1.8M
 * Total Submissions: 4.1M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
 *
 * 请你将两个数相加，并以相同形式返回一个表示和的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：l1 = [2,4,3], l2 = [5,6,4]
 * 输出：[7,0,8]
 * 解释：342 + 465 = 807.
 *
 *
 * 示例 2：
 *
 *
 * 输入：l1 = [0], l2 = [0]
 * 输出：[0]
 *
 *
 * 示例 3：
 *
 *
 * 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
 * 输出：[8,9,9,9,0,0,0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每个链表中的节点数在范围 [1, 100] 内
 * 0
 * 题目数据保证列表表示的数字不含前导零
 *
 *
 */

// @lc code=start

package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// AddTwoNumbersListNode 链表结构
type AddTwoNumbersListNode struct {
	Val  int
	Next *AddTwoNumbersListNode
}

// AddTwoNumbers 求两数之和
func AddTwoNumbers(l1 *AddTwoNumbersListNode, l2 *AddTwoNumbersListNode) *AddTwoNumbersListNode {
	carry := 0
	result := []int{}

	for {
		i, j := 0, 0
		if l1 != nil {
			i = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			j = l2.Val
			l2 = l2.Next
		}
		sum := i + j + carry

		carry = sum / 10
		result = append(result, sum%10)

		if l1 == nil && l2 == nil {
			// 两个链表都空时跳出循环
			break
		}
	}
	if carry == 1 {
		result = append(result, 1)
	}
	return SliceToLink(result)
}

// SliceToLink 将切片转为链表
func SliceToLink(arr []int) *AddTwoNumbersListNode {
	var head *AddTwoNumbersListNode
	for i := len(arr) - 1; i >= 0; i-- {
		node := &AddTwoNumbersListNode{
			Val:  arr[i],
			Next: head,
		}

		head = node
	}
	return head
}

// 返回链表的字符串形式
func (link *AddTwoNumbersListNode) toString() string {
	str := ""
	for ; link != nil; link = link.Next {
		str += fmt.Sprintf("=>%d", link.Val)
	}
	return str
}

func (link *AddTwoNumbersListNode) Compare(target *AddTwoNumbersListNode) bool {
	self := link
	targetLink := target
	for ; self != nil; self = self.Next {
		// self 未空， targetLink已空的情况
		if targetLink == nil {
			return false
		}
		if self.Val != targetLink.Val {
			return false
		}
		targetLink = targetLink.Next
	}

	// targetLink 非空时 target 比 self 长，返回 false, 否则他们长度相同，在前面已比较过数值相同，所以两个链表相等
	return targetLink == nil

}

// @lc code=end
