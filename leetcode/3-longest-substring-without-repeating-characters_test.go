package leetcode

import (
	"testing"
)

var p3Cases = []struct {
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

func p3Base(t *testing.T, f func(string) int) {

	for _, c := range p3Cases {
		len := f(c.s)
		if len != c.len {
			t.Errorf("%s %s 最长重复字串长度应为 %d, but get %d", ballotX, c.s, c.len, len)
		} else {
			t.Logf("%s string: %s,\tmaxSubLen: %d", checkMark, c.s, c.len)
		}
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("v1", func(t *testing.T) { p3Base(t, p3v1) })
	t.Run("v2", func(t *testing.T) { p3Base(t, p3v2) })
	t.Run("v3", func(t *testing.T) { p3Base(t, p3v3) })
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	b.ReportAllocs()
	s := p3Cases[0]
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(s.s)
	}
}
