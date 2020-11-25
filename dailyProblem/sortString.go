package dailyProblem

//https://leetcode-cn.com/problems/increasing-decreasing-string/
func sortString(s string) string {
	res := ""
	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
	}
	for len(res) != len(s) {
		//选小的
		for i := 0; i < len(counter); i++ {
			if counter[i] > 0 {
				res += string(rune(i + 'a'))
				counter[i]--
			}
		}
		//选大的
		for i := len(counter) - 1; i >= 0; i-- {
			if counter[i] > 0 {
				res += string(rune(i + 'a'))
				counter[i]--
			}
		}
	}
	return res
}
