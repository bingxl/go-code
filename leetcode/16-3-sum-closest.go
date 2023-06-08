/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode.cn/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (44.86%)
 * Likes:    1400
 * Dislikes: 0
 * Total Accepted:    470.5K
 * Total Submissions: 1M
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
 *
 * 返回这三个数的和。
 *
 * 假定每组输入只存在恰好一个解。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,0,0], target = 1
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 1000
 * -1000 <= nums[i] <= 1000
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start

package leetcode

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	var abs = func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	// 存放当前最短距离的三个值， [][2]int, [][0]为距离， [][1] 为距离对应的切片值
	minSum := nums[0] + nums[1] + nums[2]
	minDistance := abs(target - minSum)

	for left := 0; left < len(nums); left++ {
		for mid := left + 1; mid < len(nums); mid++ {
			for right := mid + 1; right < len(nums); right++ {

				sum := nums[left] + nums[mid] + nums[right]
				distance := abs(target - sum)

				if distance <= minDistance {
					minDistance = distance
					minSum = sum
				}

				// fmt.Printf("current: left:%d, mid: %d, right: %d, minSum: %f, mindistance: %f \n", left, mid, right, minSum, mindistance)
			}
		}
	}

	return minSum
}

// @lc code=end
