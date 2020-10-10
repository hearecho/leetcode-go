package ImplementStrstr

import "strings"

/**
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
*/
//双指针
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			//遍历needle结束全部字符相等，则i就是起始位置
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			//存在不相等的字符,则直接打破循环
			if haystack[i+j] != needle[j] {
				break
			}
		}
	}
}

//使用equal
func strStr_1(haystack string, needle string) int {
	l := len(needle)
	for i := 0; i < len(haystack)-l; i++ {
		if strings.EqualFold(haystack[i:i+l], needle) {
			return i
		}
	}
	return -1
}

//SunDay算法
//计算偏移表  偏移表的作用是存储每一个在 模式串 中出现的字符，在 模式串 中出现的最右位置到尾部的距离 +1+1
func strStr_2(haystack string, needle string) int {
	l_pattern := len(needle)
	l_haystack := len(haystack)
	offset := make(map[uint8]int)
	//计算偏移表
	for i := 0; i < len(needle); i++ {
		offset[needle[i]] = l_pattern - i
	}
	//遍历目标串
	for i := 0; i <= l_haystack-l_pattern; {
		if haystack[i:i+l_pattern] == needle {
			return i
		} else {
			if i+l_pattern > l_haystack-1 {
				return -1
			}
			//更改i
			if index,ok := offset[haystack[i+l_pattern]];ok {
				i+=index
			} else {
				i+=l_pattern+1
			}
		}
	}
	return -1
}
