package LongestValidParentheses

func longestValidParentheses(s string) int {
	//使用栈  不行
	n := len(s)
	dp := make([]int, n)
	// 使用动态规划
	res := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			//两种情况
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		res = max(res, dp[i])

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
