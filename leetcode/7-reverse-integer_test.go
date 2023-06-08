package leetcode

import "testing"

func TestP7Reverse(t *testing.T) {
	var testCases = [][]int{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{1534236469, 0}, // 益处32位则返回0
	}
	methods := map[string]func(int) int{
		"p7Reverse":   p7Reverse,
		"p7ReverseV2": p7ReverseV2,
	}
	for name, method := range methods {
		t.Run("测试："+name, func(t *testing.T) {
			for _, v := range testCases {
				testResult := method(v[0])
				if testResult == v[1] {
					t.Logf(rightSymbol)
				} else {
					t.Errorf("%s expect %d, but get: %d", errorSymbol, v[1], testResult)
				}
			}
		})
	}

}
