package Permute2

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)
	backtrack(nums, &used, &res, make([]int, 0))
	return res
}

func backtrack(nums []int, used *[]bool, res *[][]int, curArr []int) {
	if len(curArr) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, curArr)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
			continue
		}
		if !(*used)[i] {
			(*used)[i] = true
			curArr = append(curArr, nums[i])
			backtrack(nums, used, res, curArr)
			curArr = curArr[:len(curArr)-1]
			(*used)[i] = false
		}
	}
}
