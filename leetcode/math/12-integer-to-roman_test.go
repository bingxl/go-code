package math

import "testing"

func TestP12(t *testing.T) {
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

	t.Run("p12V1", func(t *testing.T) {
		for _, c := range p12Cases {
			output := intToRoman(c.value)
			if output != c.symbol {
				t.Errorf("%s, string: %v, except: %v, but got:%v", "error", c.value, c.symbol, output)
			} else {
				t.Logf("success")
			}
		}
	})
}
