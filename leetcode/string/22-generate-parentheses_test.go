package string

import (
	"slices"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {

	tests := []struct {
		name string
		args int
		want []string
	}{
		// TODO: Add test cases.
		{"Testcase 1:", 1, []string{"()"}},
		{"Testcase 2:", 2, []string{"()()", "(())"}},
		{"Testcase 3:", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.args)
			slices.Sort(got)
			slices.Sort(tt.want)
			if !slices.Equal(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
