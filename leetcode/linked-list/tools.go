package linkedList

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
