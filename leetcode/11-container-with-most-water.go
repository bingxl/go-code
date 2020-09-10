package leetcode

import (
	"fmt"
	"math"
)

// 暴力破解, 两次循环, O(n^2)
func p11v1(height []int) int {
	maxArea := 0.0

	for i := 0; (i + 1) < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := math.Min(float64(height[i]), float64(height[j])) * float64(j - i);
			
			if maxArea < area {
				maxArea = area
			}
		}
	}
	return int(maxArea)
}

// 从两侧搜寻
func p11v2(height []int) int {
	maxArea := 0

	i,j := 0, len(height) - 1

	for i != j {
		min := i
		if height[i] > height[j] {
			// j位置元素小则j左移
			min = j
			j--
		} else {
			// i位置元素小,则i右移
			i++
		}
		// j-i+1是因为i/j的自增自减在前面
		area := height[min] * (j-i+1)
		if maxArea < area {
			maxArea = area
		}
		fmt.Printf("[%v, %v], max: %v, curMax: %v \n", i, j, maxArea, area);
	
	}
	fmt.Printf("################end one case###########")

	return maxArea
}