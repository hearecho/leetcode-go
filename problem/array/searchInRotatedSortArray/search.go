package searchInRotatedSortArray

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
*/
//二分法
func search(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l <= h {
		mid := l + (h-l)/2
		if nums[mid] == target {
			return mid
		}
		//中点如果在前面的数值大则nums[mid]>nums[l]
		//如果在后面的数值小，则nums[mid]<nums[h]
		if nums[mid] > nums[l] {
			//中点在数值大的部分
			if nums[mid] > target && nums[l] <= target {
				//递增段如果 nums[mid] 大于target，并且target比nums[l]大
				h = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[h] {
			//中点在数值小的部分
			if nums[mid] < target && target <= nums[h] {
				l = mid + 1
			} else {
				h = mid - 1
			}
		} else {
			//边界相等的情况
			if nums[l] == nums[mid] {
				l++
			}
			if nums[h] == nums[mid] {
				h--
			}
		}
	}
	return -1
}
