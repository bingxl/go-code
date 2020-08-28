package leetcode

// isMatch 判断给定得字符串是否匹配给定得正则表达式(正则表达式中仅有.*)
func p10(s string, p string) bool {
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
				if p10(string(sRu[i:]), string(pRu[j+1:])) {
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
