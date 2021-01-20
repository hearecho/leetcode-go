package dailyProblem

import "sort"

func maximumProduct(nums []int) int {
	//奇数相乘 如果存在负数则比较负数 和 正数进行比较
	sort.Ints(nums)
	h := len(nums) - 1
	lowsum := nums[0] * nums[1]
	highsum := nums[h-2] * nums[h-1]
	if nums[h] >= 0 {
		if lowsum > highsum {
			return nums[h] * lowsum
		} else {
			return nums[h] * highsum
		}
	} else {
		return nums[h] * highsum
	}
}
