package cuttingRope

import "math"

func cuttingRope(n int) int {
	// 尽可能的将段分配的均匀  由于段最小长为2
	// 最优的段长为3(经过数学推算)2应该是次优 1是最差的 不能出现1
	// 所以现在就是 n小于等于3的情况下有一个大小为1的段 即返回n-1
	if n <= 3 {
		return n - 1
	}
	// 当n大于3的时间，计算余数
	a, b := n/3, n%3
	if b == 0 {
		// 余数为0 则证明可以被3整除
		return int(math.Pow(3, float64(a)))
	}
	if b == 1 {
		return int(math.Pow(3, float64(a-1)) * 4)
	}
	// 余数为2的情况
	return int(math.Pow(3, float64(a))) * 2
}

func cuttingRopeDP(n int) int {
	// 使用动态规划
	// dp[i]表示长度为i的时候最大乘积
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		//  由于分段长度最小为 1
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
