package dailyProblem

func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func maxProfit2(prices []int, fee int) int {
	n := len(prices)
	// dp[i][0] 表示第i天 手里没有股票的最大收益，而dp[i][1]表示第i天手里有股票的最大收益
	//dp[i][0] = max(dp[i-1][0],dp[i-1][1] + prices[i] - fee)
	//想持有股票 就要买
	//dp[i][1] = max(dp[i-1][1],dp[i-1][0] - prices[i])
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
