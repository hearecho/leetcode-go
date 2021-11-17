package DicesProbability

func dicesProbability(n int) []float64 {
	dp := make([]float64, 6)
	for i := 0; i < 6; i++ {
		dp[i] = 1.0 / 6.0
	}
	//n个骰子 则由 6n - n +1个数字
	for i := 2; i < n+1; i++ {
		tmp := make([]float64, 5*i+1)
		for j := 0; j < len(dp); j++ {
			for k := 0; k < 6; k++ {
				// 计算对下一个数据的贡献
				tmp[j+k] += dp[j] / 6.0
			}
		}
		dp = tmp
	}
	return dp
}
