package MaxProfit

import "math"

func maxProfit(prices []int) int {
	// dp[i]表示在前i天能够获取的最大利益
	// cost为前面i天最小的买入价格
	cost, profit := math.MaxInt64, 0
	for _, price := range prices {
		cost = min(cost, price)
		profit = max(profit, price-cost)
	}
	return profit

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxProfitViolent(prices []int) int {
	// 暴力方法
	res := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > prices[i] {
				res = max(res, prices[j]-prices[i])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
