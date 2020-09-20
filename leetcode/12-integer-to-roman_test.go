package leetcode

import "testing"

var p12Cases = []struct {
	symbol string
	value  int
}{
	{"I", 1},
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func p12Base(t *testing.T, f func(int) string) {
	for _, c := range p12Cases {
		output := f(c.value)
		if output != c.symbol {
			t.Errorf("%s, string: %v, except: %v, but got:%v", ballotX, c.value, c.symbol, output)
		} else {
			t.Logf(checkMark)
		}
	}
}

func TestP12(t *testing.T) {
	t.Run("p12V1", func(t *testing.T) { p12Base(t, p12v1) })
}
