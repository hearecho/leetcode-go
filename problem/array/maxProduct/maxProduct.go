package maxProduct

func maxProduct(nums []int) int {
	// 连续子数组
	// dp[i]表示以i为结尾的连续数组的最大和
	// 需要是连续数组 所以 dp[i]只和dp[i-1]有关系
	// 但是还有一个正负号问题 因为负负得正
	// 而负数与正数相乘则会出现一个最小值
	// 所以我们记录一下当前的最小值以及最大值，而遇到负数的时候让当前值和最小值相乘得到最大值
	if len(nums) == 1 {
		return nums[0]
	}
	CurMax, curMin, res := 1, 1, nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			// 对于负数需要改变最大值和最小值
			CurMax, curMin = curMin, CurMax
		}
		CurMax = max(nums[i], CurMax*nums[i])
		curMin = min(nums[i], curMin*nums[i])
		res = max(res, CurMax)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
