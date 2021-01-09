package dailyProblem

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
//限制 只能进行两次交易
//手里持有股票只能卖，额没有持有股票只能买

func maxProfit3(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2

}
