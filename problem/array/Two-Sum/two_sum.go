package Two_Sum

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，
并返回他们的数组下标。你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
*/

//暴力解法，时间复杂度O(n2)，空间复杂度O(1)
func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	l := len(nums)
	//code
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
			}
		}
	}
	return res
}

//针对性优化，使用空间来降低时间复杂度
func twoSum_optimize(nums []int, target int) []int {
	dict := make(map[int]int)
	l := len(nums)
	//code
	//使用字典存储索引和值
	for i := 0; i < l; i++ {
		remain := target - nums[i]
		if _,ok:=dict[remain];ok{
			return []int{dict[remain],i}
		}
		dict[nums[i]] = i
	}
	return []int{}
}
