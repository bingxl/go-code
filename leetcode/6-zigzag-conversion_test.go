package leetcode

import "testing"

func TestConvert(t *testing.T) {
	var testCases = []struct {
		sou    string
		rows   int
		result string
	}{
		{"A", 1, "A"},
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	}
	for _, ca := range testCases {
		re := convert(ca.sou, ca.rows)
		if re == ca.result {
			t.Logf(rightSymbol)
		} else {
			t.Errorf("%s 函数返回:%s,但期望:%s", errorSymbol, re, ca.result)
		}
	}
}
