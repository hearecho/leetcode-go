package TranslateNum

import "strconv"

func translateNum(num int) int {
	// dp[i]表示以第i个数字结尾的表示数量
	// 而dp[i]只和dp[i-1]以及dp[i-2]相关。
	// 即 dp[i] = dp[i-1] + dp[i-2] 当然受dp[i-2]影响的前提是 i>=2 且2=>s[i-1]>0并且 s[i-1]s[i] < 26 就是说在合理的范围内部
	s := strconv.Itoa(num)
	if len(s) == 1 {
		return 1
	}
	dp := make([]int, len(s))
	dp[0] = 1
	if s[0] == '0' || int(s[0]-'0')*10+int(s[1]-'0') > 25 {
		dp[1] = 1
	} else {
		dp[1] = 2
	}
	for i := 2; i < len(s); i++ {
		temp := int(s[i-1]-'0')*10 + int(s[i]-'0')
		if temp > 25 || s[i-1] == '0' {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[len(s)-1]
}
