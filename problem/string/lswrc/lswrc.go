package lswrc

import "strings"

//无重复最长子串
//子串必须是连续的
func lengthOfLongestSubstring(s string) int {
	//滑动窗口
	right,left := 0,0
	max := 0
	for right < len(s) {
		tempString := s[left:right]
		currentChar := s[right:right+1]
		if index := strings.Index(tempString,currentChar);index == -1{
			//不存在该字符,则更新最大长度
			max = maxNum(max,right-left+1)
			right++
		} else {
			left = index+1+len(s[:left])
		}
	}
	return max
}

func optimaize_lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 扩展 ASCII 码的位图表示（BitSet），共有 256 位
	var bitSet [256]uint8
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && bitSet[s[right]] == 0 {
			// 标记对应的 ASCII 码为 1
			bitSet[s[right]] = 1
			right++
		} else {
			// 标记对应的 ASCII 码为 0
			bitSet[s[left]] = 0
			left++
		}
		result = maxNum(result, right-left)
	}
	return result
}
func maxNum(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
