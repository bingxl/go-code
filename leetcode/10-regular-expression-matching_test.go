package leetcode

import "testing"
var p10Cases = []struct {
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
		{"ab", "a*", false},
		{"aab", "c*a*b", true},
	}
func p10Base(t *testing.T, f func(string, string)(bool)) {
	for _, c := range p10Cases {
		r := f(c.s, c.p)
		if r != c.result {
			t.Errorf("%s, string: %s, except: %v, but got:%v", ballotX, c.s, c.result, r)
		} else {
			t.Logf(checkMark)
		}
	}
}

func TestP10(t *testing.T) {
	// t.Run("p10badResolve", func(t *testing.T) {p10Base(t, p10badResolve)})
	t.Run("p10v1", func(t *testing.T) {p10Base(t, p10V1)})
}
