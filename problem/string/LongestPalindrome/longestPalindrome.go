package LongestPalindrome

func longestPalindrome(s string) string {
	maxString := ""
	if len(s) < 2 {
		return s
	}
	// 偶数类型与奇数类型的区别就是起始字符串个数不一样
	//先讨论偶数类型
	for i := 0; i < len(s)-1; i++ {
		//需要左右两端，
		if s[i] != s[i+1] {
			continue
		}
		start := s[i : i+2]
		//判断一下start
		l, r := i-1, i+2
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				if len(maxString) < len(start) {
					maxString = start
				}
				break
			}
			start = s[l:l+1] + start + s[r:r+1]
			l--
			r++
		}
		if len(maxString) < len(start) {
			maxString = start
		}
	}
	//奇数类型
	for i := 0; i < len(s); i++ {
		//需要左右两端，
		start := s[i : i+1]
		l, r := i-1, i+1
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				if len(maxString) < len(start) {
					maxString = start
				}
				break
			}
			start = s[l:l+1] + start + s[r:r+1]
			l--
			r++
		}
		if len(maxString) < len(start) {
			maxString = start
		}
	}
	return maxString
}
