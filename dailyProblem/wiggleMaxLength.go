package dailyProblem

//https://leetcode-cn.com/problems/wiggle-subsequence/

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	down, up := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}
