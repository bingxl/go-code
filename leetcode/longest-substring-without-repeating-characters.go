package leetcode

import (
	"math"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var substr string
	var maxLen float64 = 0
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if strings.Contains(substr, char) {
			maxLen = math.Max(maxLen, float64(len(substr)))
			i -= len(substr)
			substr = ""
		} else {
			substr = substr + char
		}
	}

	return int(math.Max(maxLen, float64(len(substr))))
}

func p3v1(s string) int {
	maxLenth := 0.0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if p3v1IsUnique(s, i, j) {
				maxLenth = math.Max(maxLenth, float64(j-i))
			}
		}
	}
	return int(maxLenth)
}

func p3v1IsUnique(s string, start, end int) bool {
	store := make(map[string]int)
	for i := start; i < end; i++ {
		char := string(s[i])
		if store[char] == 1 {
			return false
		}
		store[char] = 1
	}
	return true
}

func p3v2(s string) int {
	n := len(s)
	hash := make(map[string]bool)
	i, j, ans := 0, 0, 0.0
	for i < n && j < n {
		char := string(s[j])
		// fmt.Println(i, j)
		if hash[char] != true {
			hash[char] = true
			j++
			ans = math.Max(ans, float64(j-i))
		} else {
			hash[string(s[i])] = false
			i++
		}
	}
	return int(ans)
}

func p3v3(s string) int {
	n := len(s)
	hash := make(map[string]float64)
	ans := 0.0
	for i, j := 0, 0; j < n; j++ {
		char := string(s[j])
		if v, ok := hash[char]; ok {
			i = int(math.Max(v, float64(i)))
		}
		ans = math.Max(ans, float64(j-i+1))
		hash[char] = float64(j + 1)
	}
	return int(ans)
}
