/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode.cn/problems/two-sum/description/
 *
 * algorithms
 * Easy (52.87%)
 * Likes:    17130
 * Dislikes: 0
 * Total Accepted:    4.5M
 * Total Submissions: 8.5M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
 *
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
 *
 * 你可以按任意顺序返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,7,11,15], target = 9
 * 输出：[0,1]
 * 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,4], target = 6
 * 输出：[1,2]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [3,3], target = 6
 * 输出：[0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 * 只会存在一个有效答案
 *
 *
 *
 *
 * 进阶：你可以想出一个时间复杂度小于 O(n^2) 的算法吗？
 *
 */

// @lc code=start

package leetcode

// TwoSum get the two index which sum in target
//	暴力破解法
// @param	nums	[]int	输入数组
// @param	target	int		和
// @return []int 返回符合条件的一对索引
func TwoSum(nums []int, target int) []int {
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

// TwoSum get the two index which sum in target
//	使用 hash 表查找
// @param	nums	[]int	输入数组
// @param	target	int		和
// @return []int 返回符合条件的一对索引

func TwoSumV1(nums []int, target int) []int {
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

// @lc code=end
