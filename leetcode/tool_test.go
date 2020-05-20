package leetcode

import "testing"

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
			t.Logf(checkMark)
		} else {
			t.Errorf("%s, expect: %s, but get: %s", ballotX, v.output, re)
		}
	}
}
