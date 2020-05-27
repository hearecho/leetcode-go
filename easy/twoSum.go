package main

func TwoSum(nums []int, target int) []int {
	result := make([]int,2)
	i := 0
	flag := false
	for ;i<len(nums);i++ {
		result[0] = i
		y := target - nums[i]
		for j:=i+1;j<len(nums);j++ {
			if nums[j] == y {
				result[1] = j
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	return result
}
