package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {

	testCases := []struct {
		firstNum  []int
		secondNum []int
		expect    []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
	}

	t.Run("2. Add Two Numbers Test", func(t *testing.T) {

		for _, testCase := range testCases {
			result := AddTwoNumbers(SliceToLink(testCase.firstNum), SliceToLink(testCase.secondNum))

			expectLink := SliceToLink(testCase.expect)
			if result.Compare(expectLink) {
				t.Log(testCase, "   ", rightSymbol)
			} else {
				t.Error("expect: ", expectLink.toString(), " but get: ", result.toString(), "     ", errorSymbol)
			}
		}
	})
}
