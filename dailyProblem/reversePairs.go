package dailyProblem

//https://leetcode-cn.com/problems/reverse-pairs/

func reversePairs(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if 2*nums[j] < nums[i] {
				res++
			}
		}
	}
	return res
}
