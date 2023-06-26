package main

import "testing"

func TestThreeSumClosest(t *testing.T) {
	testCases := []struct {
		inputs []int
		target int
		expect int
	}{
		{inputs: []int{-1, 2, 1, -4}, target: 1, expect: 2},
		{inputs: []int{4, 0, 5, -5, 3, 3, 0, -4, -5}, target: -2, expect: -2},
		{inputs: []int{1, 1, 1, 1}, target: -100, expect: 3},
		{inputs: []int{1, 1, 1, 0}, target: -100, expect: 2},
		{inputs: []int{0, 0, 0}, target: 1, expect: 0},
	}

	testMethods := map[string]func([]int, int) int{
		"ThreeSumClosest": threeSumClosest,
	}

	t.Run("最接近目标的三个数值之和的测试", func(t *testing.T) {
		for _, tc := range testCases {
			for methodName, method := range testMethods {
				t.Run(methodName, func(t *testing.T) {
					result := method(tc.inputs, tc.target)
					if result != tc.expect {
						t.Errorf("期望得到：%d，实际得到：%d", tc.expect, result)
					}
				})
			}
		}
	})
}
