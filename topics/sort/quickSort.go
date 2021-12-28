package sort

// 快排
// 快排的重点在于选择哨兵以及划分

func quickSort(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, l, h int) {
	i := partion(nums, l, h)
	sort(nums, i+1, h)
	sort(nums, l, i-1)
}

// 分割数组的几种写法
func partion(nums []int, l, h int) int {
	// 第一种 使用双指针从两端进行遍历
	// 哨兵的话使用 nums[l]
	i, j := l, h
	watch := nums[l]
	for i < j {
		for i < j && nums[j] >= watch {
			j--
		}
		for i < j && nums[i] <= watch {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[l] = nums[l], nums[i]
	return i
}
