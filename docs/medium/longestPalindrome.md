# [5. 最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)

> 给定一个字符串 `s`，找到 `s` 中最长的回文子串。你可以假设 `s` 的最大长度为 1000。

### 解题思路

> 由于回文字符串必定是对称的，并且分为偶数和奇数。并且在原来的回文字符串两端增加字符串仍然是回文字符串。

```go
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
```

