/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (47.39%)
 * Likes:    2823
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 2.8M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given the head of a linked list, remove the n^th node from the end of the
 * list and return its head.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], n = 2
 * Output: [1,2,3,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1], n = 1
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1,2], n = 1
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is sz.
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
 *
 *
 *
 * Follow up: Could you do this in one pass?
 *
 */

// @lc code=start

package linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listLen := 0
	for cur := head; cur != nil; cur = cur.Next {
		listLen += 1
	}

	if listLen == 1 {
		return nil
	}

	index := 0
	pre := head
	removeIndex := listLen - n
	// 处理移除元素在链表头的情况
	if removeIndex == 0 {
		return head.Next
	}

	for cur := head; cur != nil; cur = cur.Next {
		if index == removeIndex {
			pre.Next = cur.Next
			cur.Next = nil
			break
		}
		pre = cur
		index++
	}
	return head
}

// 添加哑结点
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	listLen := 0
	for cur := head; cur != nil; cur = cur.Next {
		listLen++
	}
	// 哑结点
	dummy := &ListNode{0, head}

	cur := dummy
	for i := 0; i < listLen-n; i++ {
		cur = cur.Next
	}
	tmp := cur.Next
	cur.Next = cur.Next.Next
	tmp.Next = nil

	return dummy.Next

}

// @lc code=end
