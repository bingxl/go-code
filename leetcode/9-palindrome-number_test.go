package main

import "testing"

func TestP9(t *testing.T) {
	var p9Cases = []struct {
		input  int
		result bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{0, true},
		{3, true},
		{1001, true},
	}

	for _, ca := range p9Cases {
		re := isPalindrome(ca.input)
		if re == ca.result {
			t.Logf(rightSymbol)
		} else {
			t.Errorf("%s, except: %v, but get: %v", errorSymbol, ca.result, re)
		}
	}
}
