package dailyProblem

import "strconv"

//https://leetcode-cn.com/problems/monotone-increasing-digits/

func monotoneIncreasingDigits(N int) int {
	s := []rune(strconv.Itoa(N))
	i := 1
	//找到第一个递减序列
	for i < len(s) && s[i] >= s[i-1] {
		i++
	}
	// 如果全部递增则直接返回
	if i < len(s) {
		for i > 0 && s[i] < s[i-1] {
			s[i-1]--
			i--
		}
		for i++; i < len(s); i++ {
			s[i] = '9'
		}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
}
