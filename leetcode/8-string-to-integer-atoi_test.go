package main

import "testing"

func TestMyAtoi(t *testing.T) {
	var testCases = []struct {
		s string
		i int
	}{
		{"   123abd", 123},
		{"  -123", -123},
		{"+", 0},
		{"  ", 0},
		{"345a", 345},
		{"+12", 12},
		{"  w12a", 0},
		{"-91283472332", -2147483648},
		{"91283472332", 2147483647},
	}

	t.Run("P8 myAtoi 测试", func(t *testing.T) {
		for _, v := range testCases {
			r := myAtoi(v.s)
			if r == v.i {
				t.Logf(rightSymbol)
			} else {
				t.Errorf("%s, expect:%d, but get: %d", errorSymbol, v.i, r)
			}
		}
	})
}
