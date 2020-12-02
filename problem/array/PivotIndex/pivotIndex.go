package PivotIndex

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	left := 0
	for i := 0; i < len(nums); i++ {
		// right = sum-num[i]-left
		if left == sum-nums[i]-left {
			return i
		}
		left += nums[i]
	}
	return -1
}
