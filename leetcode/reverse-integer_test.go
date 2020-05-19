package leetcode

import "testing"

var p7Cases = [][]int{
	{123, 321},
	{-123, -321},
	{120, 21},
	{0, 0},
	{1534236469, 0}, // 益处32位则返回0
}

func p7Base(t *testing.T, f func(int) int) {
	for _, v := range p7Cases {
		testResult := f(v[0])
		if testResult == v[1] {
			t.Logf(checkMark)
		} else {
			t.Errorf("%s expect %d, but get: %d", ballotX, v[1], testResult)
		}
	}
}
func TestP7(t *testing.T) {
	t.Run("v1", func(t *testing.T) { p7Base(t, p7v1) })
	t.Run("v2", func(t *testing.T) { p7Base(t, p7v2) })
}
