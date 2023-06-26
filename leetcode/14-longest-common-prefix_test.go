package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	var p14Cases = []struct {
		strings []string
		common  string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"", "123", "1"}, ""},
		{[]string{}, ""},
	}

	methods := map[string]func([]string) string{
		"longestCommonPrefixV1": longestCommonPrefixV1,
		"longestCommonPrefixV2": longestCommonPrefixV2,
	}

	for name, method := range methods {
		t.Run(name, func(t *testing.T) {
			for _, c := range p14Cases {
				result := method(c.strings)
				if result != c.common {
					t.Errorf("%v, except: %s, but got: %s", errorSymbol, c.common, result)
				} else {
					t.Logf(rightSymbol)
				}
			}
		})
	}
}
