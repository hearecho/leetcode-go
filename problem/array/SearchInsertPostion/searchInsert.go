package SearchInsertPostion

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
*/
//已经排序，且无重复数组
//二分法，找到
func searchInsert(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l <= h {
		mid := l + (h-l)/2
		if nums[mid] >= target {
			h = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] >= target {
				return mid + 1
			}
			l = l + 1
		}
	}
	return 0
}
