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

	methods := map[string]func(*ListNode, int) *ListNode{
		"removeNthFromEnd":  removeNthFromEnd,
		"removeNthFromEnd1": removeNthFromEnd1,
	}
	for name, method := range methods {
		for _, tt := range tests {
			t.Run(name+"--"+tt.name, func(t *testing.T) {
				got := list2Slice(method(slice2List(tt.args.list), tt.args.n))
				if !slices.Equal(got, tt.want) {
					t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
				}
			})
		}
	}

}
