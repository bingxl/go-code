package leetcode

func p13V1(s string) int {
	romans := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	strLen := len(s)
	result := 0
	for i := 0; i < strLen; i++ {
		if (i+1) < strLen && romans[s[i]] < romans[s[i+1]] {
			result += romans[s[i+1]] - romans[s[i]]
			// 处理两个字符,所以需要多移动一位字符
			i++
		} else {
			result += romans[s[i]]
		}
	}
	return result
}
