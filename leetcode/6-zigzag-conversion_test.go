package leetcode

import "testing"

type p6CaseT struct {
	sou    string
	rows   int
	result string
}

var p6Cases = []p6CaseT{
	{"A", 1, "A"},
	{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
}

func TestP6v1(t *testing.T) {
	for _, ca := range p6Cases {
		re := p6v1(ca.sou, ca.rows)
		if re == ca.result {
			t.Logf(checkMark)
		} else {
			t.Errorf("%s 函数返回:%s,但期望:%s", ballotX, re, ca.result)
		}
	}
}
