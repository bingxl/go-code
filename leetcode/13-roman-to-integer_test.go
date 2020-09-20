package leetcode

import "testing"

func TestP13(t *testing.T) {
	for _, c := range p12Cases {
		out := p13V1(c.symbol)
		if out != c.value {
			t.Errorf("roman: %s, expect: %d, but got: %d", c.symbol, c.value, out)
		} else {
			t.Log(checkMark)
		}
	}
}
