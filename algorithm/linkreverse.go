package algorithm

import "fmt"

// LNode 链表节点类型
type LNode struct {
	Val  interface{}
	Next *LNode
}

// Reverse 带头节点的链表逆序
func (node *LNode) Reverse() {
	if node == nil || node.Next == nil {
		return
	}

	var pre, cur *LNode
	next := node.Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

// CreateLink 创建链表
func CreateLink(nums []interface{}) *LNode {
	var head *LNode
	for i := len(nums) - 1; i >= 0; i-- {
		node := &LNode{
			Val:  nums[i],
			Next: head,
		}

		head = node
	}
	return head
}

// Print 打印链表
func (node *LNode) Print() {
	fmt.Println(node.ToString())
}

// ToString 获取两链表字符串形式
func (node *LNode) ToString() string {
	cur := node.Next
	str := ""
	for cur != nil {
		str += "=>" + interface2String(cur.Val)
		cur = cur.Next
	}
	return str
}
