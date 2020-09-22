package leetcode

import (
	"strings"
)

func p14V1(strs []string) string {
	if (len(strs)) == 0 {
		return ""
	}
	minLen := len(strs[0])
	commonPrefix := strings.Builder{}

	for _, s := range strs {
		if minLen > len(s) {
			minLen = len(s)
		}
	}

	for i := 0; i < minLen; i++ {
		m := make(map[byte]int)
		var commonByte byte
		for _, s := range strs {
			commonByte = s[i]
			m[commonByte]++
		}
		if len(m) != 1 {
			break
		} else {
			commonPrefix.WriteByte(commonByte)
		}
	}

	return commonPrefix.String()
}

func p14V2(strs []string) string {
	// 如果为空,返回"", 如果只有一个字符串则返回这个字符串, 因为后面遍历比较时可以不用和strs[0]比较
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	baseStr := strs[0]

	for _, s := range strs[1:] {
		// 如果s没有对应的baseStr前缀则不断缩短baseStr
		for strings.Index(s, baseStr) != 0 {
			// 基串为空,直接返回
			if baseStr == "" {
				return baseStr
			}
			baseStr = baseStr[:len(baseStr)-1]
		}
	}
	return baseStr
}
