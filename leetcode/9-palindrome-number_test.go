package leetcode

import "testing"

var p9Cases = []struct {
	input  int
	result bool
}{
	{121, true},
	{-121, false},
	{10, false},
	{0, true},
	{3, true},
	{1001, true},
}

func TestP9(t *testing.T) {
	for _, ca := range p9Cases {
		re := p9v1(ca.input)
		if re == ca.result {
			t.Logf(checkMark)
		} else {
			t.Errorf("%s, except: %v, but get: %v", ballotX, ca.result, re)
		}
	}
}
