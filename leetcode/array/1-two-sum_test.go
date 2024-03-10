package array

import (
	"slices"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		expect []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, expect: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expect: []int{1, 2}},
		{nums: []int{3, 2, 4}, target: 6, expect: []int{2, 1}},
		{nums: []int{3, 3}, target: 6, expect: []int{0, 1}},
	}

	methods := map[string]func([]int, int) []int{
		"TwoSum":   TwoSum,
		"TwoSumV1": TwoSumV1,
	}

	compareSlicesWithDisorder := func(slice1, slice2 []int) bool {

		// 排序切片
		sort.Ints(slice1)
		sort.Ints(slice2)

		return slices.Equal(slice1, slice2)
	}

	t.Run("1. TwoSum 测试", func(t *testing.T) {
		for name, method := range methods {
			t.Run(name+"测试", func(t *testing.T) {
				for _, testCase := range testCases {
					result := method(testCase.nums, testCase.target)
					if !compareSlicesWithDisorder(result, testCase.expect) {
						t.Error("expect: ", testCase.expect, " but get: ", result)
					} else {
						t.Log(name, testCase, "测试通过")

					}
				}
			})
		}
	})

}
