package array

import (
	"fmt"
	"slices"
	"sort"
	"testing"
	"time"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	// 测试的函数集
	funcs := []struct {
		funName string
		fun     func([]int, int) [][]int
	}{
		{funName: "fourSum", fun: fourSum},
		{funName: "fourSum1", fun: fourSum1},
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "常规",
			args: args{nums: []int{1, 0, -1, 0, -2, 2}, target: 0},
			want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name: "ex2",
			args: args{nums: []int{2, 2, 2, 2, 2}, target: 8},
			want: [][]int{{2, 2, 2, 2}},
		},

		{
			name: "ex3",
			args: args{nums: []int{-5, 5, 4, -3, 0, 0, 4, -2}, target: 4},
			want: [][]int{{-5, 0, 4, 5}, {-3, -2, 4, 5}},
		},
		{
			name: "ex4",
			args: args{nums: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, target: 8},
			want: [][]int{{2, 2, 2, 2}},
		},
	}

	assertEqual := func(src [][]int, dst [][]int) bool {
		if len(src) != len(dst) {
			return false
		}
		for i := range len(src) {
			sort.Ints(src[i])
			sort.Ints(dst[i])
		}

		sort.Slice(src, func(i, j int) bool {
			return fmt.Sprintf("%v", src[i]) < fmt.Sprintf("%v", src[j])
		})
		sort.Slice(dst, func(i, j int) bool {
			return fmt.Sprintf("%v", dst[i]) < fmt.Sprintf("%v", dst[j])
		})
		return slices.EqualFunc(src, dst, func(e1, e2 []int) bool {
			return slices.Equal(e1, e2)
		})
	}
	for _, fun := range funcs {
		for _, tt := range tests {
			t.Run(fun.funName+"--"+tt.name, func(t *testing.T) {
				start := time.Now()
				got := fun.fun(tt.args.nums, tt.args.target)
				t.Logf("运行%s测试耗时: %v", tt.name, time.Since(start))
				if !assertEqual(got, tt.want) {
					t.Errorf("fourSum() = %v, want %v", got, tt.want)
				}
			})
		}
	}

}
