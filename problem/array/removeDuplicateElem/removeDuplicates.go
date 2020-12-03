package removeDuplicateElem

/**
双指针
*/
func removeDuplicates(nums []int) int {
	l := len(nums)
	i := 0
	j := 1
	for j < l {
		for j<l && nums[j] == nums[i] {
			j++
		}
		if j >= l {
			break
		}
		nums[i+1] = nums[j]
		i = i+1
	}
	return i+1
}
