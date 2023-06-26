package main

import "testing"

var p13Cases = []struct {
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

func TestP13(t *testing.T) {
	for _, c := range p13Cases {
		out := romanToInt(c.symbol)
		if out != c.value {
			t.Errorf("roman: %s, expect: %d, but got: %d", c.symbol, c.value, out)
		} else {
			t.Log(rightSymbol)
		}
	}
}

func BenchmarkP13(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		j := i % len(p13Cases)
		romanToInt(p13Cases[j].symbol)
	}
}
