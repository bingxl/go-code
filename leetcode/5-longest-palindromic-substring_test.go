package main

import "testing"

func TestP5v1(t *testing.T) {
	var testCases = []([]string){
		[]string{"", ""},
		[]string{" ", " "},
		[]string{"b", "b"},
		[]string{"cbbd", "bb"},
		[]string{"bb", "bb"},
		[]string{"aaabaaaa", "aaabaaa"},
		[]string{"abcb", "bcb"},
		[]string{"ccc", "ccc"},
		[]string{"cccbbbaaaaaabbbbb", "bbbaaaaaabbb"},
		[]string{"cccbbbaaaaaaabbbbb", "bbbaaaaaaabbb"},
	}
	for _, v := range testCases {
		result := longestPalindrome(v[0])
		if result == v[1] {
			t.Logf(rightSymbol)
		} else {
			t.Errorf("%s, get: %s, but should be: %s", errorSymbol, result, v[1])
		}
	}
}
