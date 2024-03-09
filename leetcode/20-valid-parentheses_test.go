package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "一对", args: args{"()"}, want: true},
		{name: "三对", args: args{"()[]{}"}, want: true},
		{name: "两个不匹配", args: args{"(]"}, want: false},
		{name: "交叉匹配", args: args{"()[([{}])]{}"}, want: true},
		{name: "交叉不匹配", args: args{"()[[({)]]{}"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
