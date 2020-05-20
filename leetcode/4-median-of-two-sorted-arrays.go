package leetcode

import "sort"

func p4v1(nums1, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	n := len(nums)
	if n%2 == 0 {
		return float64(nums[n/2]+nums[n/2-1]) / 2.0
	}
	return float64(nums[n/2])
}
