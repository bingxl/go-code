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

func BenchmarkP13(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		j := i % len(p12Cases)
		p13V1(p12Cases[j].symbol)
	}
}
