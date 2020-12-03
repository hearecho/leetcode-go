package findfirstlastIndex

/**
二分法
采用找左边界的版本的二分查找
*/
func searchRange(nums []int, target int) []int {
	left := binarySearch(nums, target)
	right := binarySearch(nums, target+1) - 1
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	} else {
		return []int{left, Max(left, right)}
	}
}

//在含有重复数组中查找元素的最左边界 则找比target大的值，即使找不到l位置的值也会大于等于target
func binarySearch(nums []int, key int) int {
	//注意h的取值
	l, h := 0, len(nums)
	for l < h {
		m := l + (h-l)/2
		if nums[m] >= key {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
