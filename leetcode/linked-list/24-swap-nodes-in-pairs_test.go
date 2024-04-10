package linkedList

import (
	"testing"
)

func Test_swapPairs(t *testing.T) {

	// Test cases
	cases := []struct {
		input  []int
		output []int
	}{
		{input: []int{1, 2, 3, 4}, output: []int{2, 1, 4, 3}},
		{input: []int{}, output: []int{}}, // 空集
		{input: []int{1}, output: []int{1}},
		{input: []int{5, 7, 9, 2, 4, 6}, output: []int{7, 5, 2, 9, 6, 4}},
		{input: []int{2, 4, 6, 8}, output: []int{4, 2, 8, 6}},
		{input: []int{1, 2, 3, 4, 5}, output: []int{2, 1, 4, 3, 5}}, // 奇数个元素类型
		// Add more test cases here
	}

	// Test each case
	for _, c := range cases {
		got := swapPairs(slice2List(c.input))
		if !isEqual(got, slice2List(c.output)) {
			t.Errorf("swapPairs(%v) == %v, want %v", c.input, got, c.output)
		}
	}

}
