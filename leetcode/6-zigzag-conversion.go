package leetcode

// 思路：使用numRows个切片来存储字符，注意正向或反向存储的问题

// 0     6      12
// 1   5 7   11 13
// 2 4   8 10   14
// 3     9      15
func p6v1(s string, numRows int) string {
	// 只需要一行时直接返回
	if numRows <= 1 {
		return s
	}
	ru := []rune(s)
	re := make([]([]rune), numRows)
	reverse := false
	for k, v := range ru {
		i := k % (numRows - 1)
		j := i
		// 第一列最后一行的值开始反向存储
		if k != 0 && i == 0 {
			reverse = !reverse
		}

		if reverse {
			j = numRows - 1 - i
		}

		re[j] = append(re[j], v)
	}

	var result string
	for _, sli := range re {
		for _, v := range sli {
			result += string(v)
		}
	}
	// fmt.Println(ru)
	// fmt.Println(re)
	return result
}
