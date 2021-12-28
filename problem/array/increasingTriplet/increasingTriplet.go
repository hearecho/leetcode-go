package increasingTriplet

import "math"

func increasingTriplet(nums []int) bool {
	// 一个简单的规律
	// 索引是递增的 所以我们只要找到一个按顺序递增的子序列即可
	// 使用动态规划 只要存在等于3的子序列可以直接返回
	// 但是这个时间复杂度实际上还是O(n2)
	if len(nums) < 3 {
		return false
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			if dp[i] >= 3 {
				return true
			}
		}
	}
	return false
}

func increasingTripletOp(nums []int) bool {
	// 时间复杂度O(n) 空间复杂度O(1)
	// 维护一个start值 和 第二个值 mid 之后出现的值 如果大于mid则直接返回true
	// 如果 小于mid大于start 则替换start值
	// 如果小于 start 则替换start
	if len(nums) < 3 {
		return false
	}
	l, mid := nums[0], math.MaxInt64
	for _, num := range nums {
		if num <= l {
			l = num
		} else if num <= mid {
			mid = num
		} else {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
