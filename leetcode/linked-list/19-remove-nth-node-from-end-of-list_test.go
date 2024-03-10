package linkedList

import (
	"slices"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		list []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "ex1",
			args: args{list: []int{1, 2, 3, 4, 5}, n: 2},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "ex2",
			args: args{list: []int{1}, n: 1},
			want: []int{},
		},
		{
			name: "ex3",
			args: args{list: []int{1, 2}, n: 1},
			want: []int{1},
		},
		{
			name: "ex4",
			args: args{list: []int{1, 2}, n: 2},
			want: []int{2},
		},
	}
	slice2List := func(ls []int) *ListNode {
		head := &ListNode{}
		cur := head
		for _, v := range ls {
			cur.Next = &ListNode{Val: v}
			cur = cur.Next
		}
		return head.Next
	}
	list2Slice := func(ls *ListNode) []int {
		result := []int{}
		for cur := ls; cur != nil; cur = cur.Next {
			result = append(result, cur.Val)
		}
		return result
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := list2Slice(removeNthFromEnd(slice2List(tt.args.list), tt.args.n))
			if !slices.Equal(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
