package leetcode

import "fmt"

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

// Print 打印链表
func (link *AddTwoNumbersListNode) Print() {
	fmt.Printf("head")
	for ; link != nil; link = link.Next {
		fmt.Printf("=>%d", link.Val)
	}
	fmt.Printf("\n")
}
