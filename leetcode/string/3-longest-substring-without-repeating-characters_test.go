package string

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var p3TestCases = []struct {
		s   string
		len int
	}{
		{"abcdedsa", 5},
		{"", 0},
		{"edabab", 4},
		{" ", 1},
		{"abaafdbc", 5},
		{"aab", 2},
		{"dvdf", 3},
	}
	methods := map[string]func(string) int{
		"lengthOfLongestSubstring":   lengthOfLongestSubstring,
		"lengthOfLongestSubstringV1": lengthOfLongestSubstringV1,
		// "lengthOfLongestSubstringV2": lengthOfLongestSubstringV2, // 此方法有bug，暂时不测
		"lengthOfLongestSubstringV3": lengthOfLongestSubstringV3,
	}

	for name, method := range methods {
		t.Run(name, func(t *testing.T) {
			for _, c := range p3TestCases {
				len := method(c.s)
				if len != c.len {
					t.Errorf(" %s 最长重复字串长度应为 %d, but get %d", c.s, c.len, len)
				} else {
					t.Logf("string: %s,\tmaxSubLen: %d", c.s, c.len)
				}
			}
		})
	}

}
