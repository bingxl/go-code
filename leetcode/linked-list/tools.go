package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func slice2List(ls []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, v := range ls {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head.Next
}
func list2Slice(ls *ListNode) []int {
	result := []int{}
	for cur := ls; cur != nil; cur = cur.Next {
		result = append(result, cur.Val)
	}
	return result
}

// 判断两个链表是否相等
func isEqual(ls1, ls2 *ListNode) bool {
	cur1, cur2 := ls1, ls2
	for ; cur1 != nil && cur2 != nil; cur1, cur2 = cur1.Next, cur2.Next {
		if cur1.Val != cur2.Val {
			return false
		}
	}
	return cur1 == nil && cur2 == nil
}
