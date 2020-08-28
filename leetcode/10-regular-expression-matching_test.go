package leetcode

import "testing"

func TestP10(t *testing.T) {
	cases := []struct {
		s      string
		p      string
		result bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"aa", "a.", true},
		{"aaa", "a.a", true},
		{"aaa", "a*aa", true},
		{"aaa", "a*", true},
		{"ab", ".*", true},
		{"ab", "*", true},
		{"aab", "c*a*b", false},
	}

	for _, c := range cases {
		r := p10(c.s, c.p)
		if r != c.result {
			t.Errorf("%s, string: %s, except: %v, but got:%v", ballotX, c.s, c.result, r)
		} else {
			t.Logf(checkMark)
		}
	}
}
