package leetcode

import "testing"

type p4CaseType struct {
	t1 []int
	t2 []int
	v  float64
}

var p4Cases = []p4CaseType{
	{[]int{1, 2, 3, 4}, []int{5, 6, 7}, 4},
	{[]int{3, 4, 5, 6, 9}, []int{0, 4, 10}, 4.5},
}

func p4Base(t *testing.T, f func([]int, []int) float64) {

	for _, suilt := range p4Cases {
		result := f(suilt.t1, suilt.t2)
		if result == suilt.v {
			t.Logf("%s", checkMark)
		} else {
			t.Errorf("%s, get %f, but should equar %f", ballotX, result, suilt.v)
		}
	}
}

func TestP4(t *testing.T) {
	t.Run("v1", func(t *testing.T) { p4Base(t, p4v1) })
}
