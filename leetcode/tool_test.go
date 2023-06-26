package main

import (
	"strings"
	"testing"
)

func TestSliceToString(t *testing.T) {
	type inputType []interface{}

	cases := []struct {
		input  inputType
		output string
	}{
		{inputType{1, 2, 3}, "->1->2->3"},
		{inputType{"a", "b", "c", 4}, "->a->b->c->4"},
	}

	for _, v := range cases {
		re := SliceToString(v.input)
		if re == v.output {
			t.Logf(rightSymbol)
		} else {
			t.Errorf("%s, expect: %s, but get: %s", errorSymbol, v.output, re)
		}
	}
}

func BenchmarkStringJoin(b *testing.B) {
	// 经过测试,如果需要多次拼接字符串使用strings.Builder{}性能更好点, 但strings.Builder在实例化时耗时
	b.Run("use plus ", func(b *testing.B) {
		s := ""
		for i := 0; i < b.N; i++ {

			s += "1"
		}
	})

	b.Run("use strings.Builder", func(b *testing.B) {
		s := strings.Builder{}
		for i := 0; i < b.N; i++ {
			s.WriteString("1")
		}
	})
}
