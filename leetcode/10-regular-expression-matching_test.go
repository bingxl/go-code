package main

import "testing"

func TestP10IsMatch(t *testing.T) {
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

	t.Run("p10IsMatch", func(t *testing.T) {
		for _, c := range p10Cases {
			r := p10IsMatch(c.s, c.p)
			if r != c.result {
				t.Errorf("%s, string: %s, except: %v, but got:%v", errorSymbol, c.s, c.result, r)
			} else {
				t.Logf(rightSymbol)
			}
		}
	})
}
