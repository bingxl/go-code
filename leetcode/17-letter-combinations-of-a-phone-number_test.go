package main

import (
	"slices"
	"sort"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{name: "两位数", args: args{digits: "23"}, want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{name: "一位数", args: args{digits: "2"}, want: []string{"a", "b", "c"}},
		{name: "一位数", args: args{digits: "7"}, want: []string{"p", "q", "r", "s"}},
		{name: "空字符串", args: args{digits: ""}, want: []string{}},
		{name: "一位数", args: args{digits: "2"}, want: []string{"a", "b", "c"}},
		// {name: "三位数位数", args: args{digits: "2"}, want: []string{"a", "b", "c"}},
	}
	// 判断两个切片中的内容是否相同, 与顺序无关
	sliceEquIgnoreSort := func(slice1, slice2 []string) bool {
		if len(slice1) != len(slice2) {
			return false
		}
		sort.Strings(slice1)
		sort.Strings(slice2)
		return slices.Equal(slice1, slice2)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.args.digits)
			if !sliceEquIgnoreSort(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			} else {
				t.Logf("success got: %v", got)
			}
		})
	}
}
