package linkedList

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		lists    [][]int
		expected []int
	}{
		{
			"empty lists",
			[][]int{},
			[]int{},
		},
		{
			"one list",
			[][]int{{1, 2, 3}},
			[]int{1, 2, 3},
		},
		{
			"two lists",
			[][]int{{1, 3, 5}, {2, 4, 6}},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"three lists",
			[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lists := make([]*ListNode, len(tt.lists))
			for i, v := range tt.lists {
				lists[i] = slice2List(v)
			}

			actual := mergeKLists(lists)
			if !isEqual(actual, slice2List(tt.expected)) {
				t.Errorf("mergeKLists() = %v, want %v", list2Slice(actual), tt.expected)
			} else {
				t.Log("PASS  " + tt.name)
			}
		})
	}
}
