package main

import (
	"fmt"
	"slices"
	"sort"
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

	equalJudge := func(expect, result [][]int) bool {
		if len(expect) != len(result) {
			return false
		}

		// 将切片每项子切片进行排序
		for i := range expect {
			sort.Ints(expect[i])
			sort.Ints(result[i])
		}

		// 将切片进行排序
		sort.Slice(expect, func(i, j int) bool {
			return fmt.Sprintf("%v", expect[i]) < fmt.Sprintf("%v", expect[j])
		})
		sort.Slice(result, func(i, j int) bool {
			return fmt.Sprintf("%v", result[i]) < fmt.Sprintf("%v", result[j])
		})

		// 比较两个切片是否相等
		return slices.EqualFunc(expect, result, func(i, j []int) bool {
			return slices.Equal(i, j)
		})
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
				if equalJudge(result, c.result) {
					t.Logf(rightSymbol)
				} else {
					t.Errorf("%s, expect: %v, but got: %v", errorSymbol, c.result, result)
				}
			}
		})
	}
}
