package dfs

//https://leetcode-cn.com/leetbook/read/queue-stack/ga4o2/

/**
使用dfs进行递归搜搜
*/

func findTargetSumWays(nums []int, S int) int {
	count := 0
	dfs(nums, 0, 0, S, &count)
	return count
}

func dfs(nums []int, index, sum int, S int, count *int) {
	if index == len(nums) {
		if sum == S {
			*count++
		}
		return
	}
	dfs(nums, index+1, sum+nums[index], S, count)
	dfs(nums, index+1, sum-nums[index], S, count)
}
