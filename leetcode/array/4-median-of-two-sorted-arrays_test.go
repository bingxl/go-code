package array

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {

	testCases := []struct {
		num1   []int
		num2   []int
		expect float64
	}{
		{[]int{1, 2, 3, 4}, []int{5, 6, 7}, 4},
		{[]int{3, 4, 5, 6, 9}, []int{0, 4, 10}, 4.5},
	}
	t.Run("4. FindMedianSortedArrays", func(t *testing.T) {
		for _, testCase := range testCases {
			result := findMedianSortedArrays(testCase.num1, testCase.num2)
			if result == testCase.expect {
				t.Log("测试用例 ", testCase)
			} else {
				t.Error("测试用例 ", testCase, " 期望值：", testCase.expect, " 获得的：", result)
			}
		}
	})
}
