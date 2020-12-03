package CanJump

func canJump(nums []int) bool {
	l := len(nums)
	for i, next_indx := 0, 0; i < l; {
		if next_indx != l-1 && nums[next_indx] == 0 {
			return false
		}
		if next_indx == l-1 {
			return true
		}
		next_indx++
		for j := 1; j <= nums[i]; j++ {
			if i+j < l && next_indx+nums[next_indx] <= nums[i+j]+i+j {
				next_indx = i + j
			}
		}
		i = next_indx
	}
	return false
}
