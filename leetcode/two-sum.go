package leetcode

// TwoSum get the two index which sum in target
//
// @param	nums	[]int	输入数组
// @param	target	int		和
// @return
func TwoSum(nums []int, target int) []int {
	// return twoForLoop(nums, target)
	return twoPassHashTable(nums, target)
}

func twoForLoop(nums []int, target int) []int {
	// time complexity O(n^2)
	// space complexity O(n)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoPassHashTable(nums []int, target int) []int {
	// Time complexity O(n)
	// Space complexity O(n)

	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = i
	}

	for j := 0; j < len(nums); j++ {
		complement := target - nums[j]
		v, ok := hash[complement]
		if ok && v != j {
			return []int{j, v}
		}
	}
	return []int{}
}

func onwPassHashTable(nums []int, target int) []int {
	// Time complexity O(n)
	// Space complexity O(n)
	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		v, ok := hash[complement]
		if ok {
			return []int{v, i}
		}
		hash[nums[i]] = i
	}
	return []int{}
}
