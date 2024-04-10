/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.cn/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (72.14%)
 * Likes:    2193
 * Dislikes: 0
 * Total Accepted:    826.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head. You
 * must solve the problem without modifying the values in the list's nodes
 * (i.e., only nodes themselves may be changed.)
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4]
 * Output: [2,1,4,3]
 *
 *
 * Example 2:
 *
 *
 * Input: head = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1]
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 100].
 * 0 <= Node.val <= 100
 *
 *
 */
package linkedList

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

	dummyNode := &ListNode{Next: head}

	cur := dummyNode

	for head != nil && head.Next != nil {

		// 链表元素交换
		cur.Next = head.Next
		head.Next = head.Next.Next
		cur.Next.Next = head

		// 指针移动, cur指针移动两次
		cur = head
		// head 指针在链表元素交换时移动过一次,此时只需要在移动一次
		head = head.Next
	}
	return dummyNode.Next
}

// @lc code=end
