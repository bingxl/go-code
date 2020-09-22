package leetcode

import "testing"

var p14Cases = []struct {
	strs   []string
	common string
}{
	{[]string{"flower", "flow", "flight"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
	{[]string{"", "123", "1"}, ""},
	{[]string{}, ""},
}

func TestP14V1(t *testing.T) {
	for _, c := range p14Cases {
		result := p14V1(c.strs)
		if result != c.common {
			t.Errorf("%v, except: %s, but got: %s", ballotX, c.common, result)
		} else {
			t.Logf(checkMark)
		}
	}
}

func TestP14V2(t *testing.T) {
	for _, c := range p14Cases {
		result := p14V2(c.strs)
		if result != c.common {
			t.Errorf("%v, expect: %s, but got: %s", ballotX, c.common, result)
		} else {
			t.Logf(checkMark)
		}
	}
}
