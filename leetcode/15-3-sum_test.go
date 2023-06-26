package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	var p15Cases = []struct {
		item   []int
		result [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{[]int{}, [][]int{}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	methods := map[string]func([]int) [][]int{
		"p15ThreeSumV1": p15ThreeSumV1,
		"p15ThreeSumV2": p15ThreeSumV2,
		"p15ThreeSumV3": p15ThreeSumV3,
	}

	for name, method := range methods {
		t.Run(name, func(t *testing.T) {
			for _, c := range p15Cases {
				result := method(c.item)

				// @TODO 判断是否相等的方法有问题
				if reflect.DeepEqual(result, c.result) {
					t.Logf(rightSymbol)
				} else {
					t.Errorf("%s, expect: %v, but got: %v", errorSymbol, c.result, result)
				}
			}
		})
	}
}
