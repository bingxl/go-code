package leetcode

import "testing"

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
