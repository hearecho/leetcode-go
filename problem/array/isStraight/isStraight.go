package isStraight

import "sort"

func isStraight(nums []int) bool {
	// 大小王可以看作是任意一个数字 所以我们需要记录0的个数
	sort.Ints(nums)
	joker := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			joker++
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			return false
		}
	}
	return nums[4]-nums[joker] < 5
}
