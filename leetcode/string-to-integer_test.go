package leetcode

import "testing"

var p8Cases = []struct {
	s string
	i int
}{
	{"   123abd", 123},
	{"  -123", -123},
	{"+", 0},
	{"  ", 0},
	{"345a", 345},
	{"+12", 12},
	{"  w12a", 0},
	{"-91283472332", -2147483648},
	{"91283472332", 2147483647},
}

func p8Base(t *testing.T, f func(string) int) {
	for _, v := range p8Cases {
		r := f(v.s)
		if r == v.i {
			t.Logf(checkMark)
		} else {
			t.Errorf("%s, expect:%d, but get: %d", ballotX, v.i, r)
		}
	}
}
func TestP8(t *testing.T) {
	t.Run("v1", func(t *testing.T) { p8Base(t, p8v1) })
}
