package MaxSubArray

func maxSubArray(nums []int) int {
	max := nums[0]
	cursum := 0
	for _, num := range nums {
		if cursum <= 0 {
			cursum = num
		} else {
			cursum += num
		}
		max = Max(max, cursum)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
