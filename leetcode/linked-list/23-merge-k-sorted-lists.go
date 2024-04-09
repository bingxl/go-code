/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.cn/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (59.17%)
 * Likes:    2801
 * Dislikes: 0
 * Total Accepted:    789.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * You are given an array of k linked-lists lists, each linked-list is sorted
 * in ascending order.
 *
 * Merge all the linked-lists into one sorted linked-list and return it.
 *
 *
 * Example 1:
 *
 *
 * Input: lists = [[1,4,5],[1,3,4],[2,6]]
 * Output: [1,1,2,3,4,4,5,6]
 * Explanation: The linked-lists are:
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * merging them into one sorted list:
 * 1->1->2->3->4->4->5->6
 *
 *
 * Example 2:
 *
 *
 * Input: lists = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: lists = [[]]
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] is sorted in ascending order.
 * The sum of lists[i].length will not exceed 10^4.
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
func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	curs := make([]*ListNode, k)
	copy(curs, lists)

	// headers为哑元, 方便后续插入
	header := &ListNode{}
	cur := header

	// 找到最小值
	minNode := func(nodes []*ListNode) (int, *ListNode) {
		var min *ListNode
		var index int
		for i, v := range nodes {
			// find min
			if v != nil && (min == nil || v.Val < min.Val) {
				min = v
				index = i
			}
		}
		// fmt.Printf("nodes: %+v, min: %v, index: %d \n", nodes, min, index)
		return index, min
	}

	for {
		index, node := minNode(curs)
		if node == nil {
			// 已循环完毕
			break
		}
		cur.Next = node
		cur = cur.Next
		curs[index] = curs[index].Next

	}
	return header.Next
}

// @lc code=end
