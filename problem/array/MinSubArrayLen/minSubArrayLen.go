package MinSubArrayLen

import "math"

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, 0
	curSum := 0
	ans := math.MaxInt32
	for r < len(nums) {
		curSum += nums[r]
		for curSum >= s {
			ans = min(r-l+1, ans)
			curSum -= nums[l]
			l++
		}
		r++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
