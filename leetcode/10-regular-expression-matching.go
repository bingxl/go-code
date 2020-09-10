package leetcode

// isMatch 判断给定得字符串是否匹配给定得正则表达式(正则表达式中仅有.*)

// 此方法是没看leetcode解法时思索的,未能正确匹配, 关键在于没善用递归
func p10badResolve(s string, p string) bool {
	sRu := []rune(s)
	pRu := []rune(p)

	i, j := 0, 0
	for j < len(pRu) {
		switch pRu[j] {
		case '.':
			{
				// s中已没有字符,则匹配失败
				if i >= len(sRu) {
					return false
				}
				i++
				j++
			}
		case '*':
			{
				if j == len(pRu)-1 {
					return true
				}
				if i >= len(sRu) {
					return false
				}
				if p10badResolve(string(sRu[i:]), string(pRu[j+1:])) {
					return true
				}
				i++
			}
		default:
			{
				if i >= len(sRu) {
					return false
				}
				if sRu[i] != pRu[j] {
					return false
				}
				i++
				j++

			}
		}
	}

	return i == len(sRu) && j == len(pRu)
}

func p10V1(s, p string) bool {
	// 正则已空时看字符串是否为空
	if len(p) == 0 {return len(s) == 0}

	// 字符串第一个和正则第一个是否匹配
	firstMatch := (len(s) != 0 && ((s[0] == p[0]) || (p[0] == '.') ))

	if len(p) >= 2 && p[1] == '*' {
		return p10V1(s, p[2:]) ||(firstMatch && p10V1(s[1:], p))
	}
	return firstMatch && p10V1(s[1:], p[1:])
}