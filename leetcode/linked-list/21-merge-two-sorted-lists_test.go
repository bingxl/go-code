package linkedList

import (
	"slices"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "both lists empty",
			args: args{
				list1: []int{},
				list2: []int{},
			},
			want: []int{},
		},
		{
			name: "list1 empty",
			args: args{
				list1: []int{},
				list2: []int{1, 3, 5},
			},
			want: []int{1, 3, 5},
		},
		{
			name: "list2 empty",
			args: args{
				list1: []int{2, 4, 6},
				list2: []int{},
			},
			want: []int{2, 4, 6},
		},
		{
			name: "lists with equal sizes",
			args: args{
				list1: []int{1, 3, 5},
				list2: []int{2, 4, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "lists with different sizes",
			args: args{
				list1: []int{1},
				list2: []int{2, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "lists with negative numbers",
			args: args{
				list1: []int{-3, 1, 4},
				list2: []int{-2, 3, 5},
			},
			want: []int{-3, -2, 1, 3, 4, 5},
		},
		{
			name: "lists with all elements of list2 smaller",
			args: args{
				list1: []int{4, 5, 6},
				list2: []int{1, 2, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "lists with duplicates",
			args: args{
				list1: []int{1, 3, 3, 5},
				list2: []int{1, 2, 3},
			},
			want: []int{1, 1, 2, 3, 3, 3, 5},
		},
	}

	testMethods := map[string]func(l1, l2 *ListNode) *ListNode{
		"mergeTwoLists":  mergeTwoLists,
		"mergeTwoLists1": mergeTwoLists1,
	}
	for name, method := range testMethods {
		for _, tt := range tests {
			t.Run(name+"--"+tt.name, func(t *testing.T) {
				got := list2Slice(method(slice2List(tt.args.list1), slice2List(tt.args.list2)))
				if !slices.Equal(got, tt.want) {
					t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func Benchmark_mergeTwoLists(b *testing.B) {
	s1 := []int{1, 3, 3, 5, 6, 7, 8, 9}
	s2 := []int{1, 3, 3, 5, 9, 11}
	for range b.N {
		mergeTwoLists(slice2List(s1), slice2List(s2))
	}
}

// 317.6 ns/op 	224 B/op	14 allocs/op
func Benchmark_mergeTwoLists1(b *testing.B) {
	s1 := []int{1, 3, 3, 5, 6, 7, 8, 9}
	s2 := []int{1, 3, 3, 5, 9, 11}
	for range b.N {
		mergeTwoLists1(slice2List(s1), slice2List(s2))
	}
}
