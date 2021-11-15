package reverseLeftWords

func reverseLeftWords(s string, n int) string {
	if n == len(s) {
		return s
	}
	// 先增加字符串长度 增加的长度为 n 之后将前面n个字符删除即可
	// 空间复杂O(n)
	res := s[n:] + s[:n]
	return res
}
