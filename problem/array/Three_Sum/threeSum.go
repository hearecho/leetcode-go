package Three_Sum

import "sort"

//最容易想到的方法，将整个数组进行排列组合，然后将三数之和的结果返回. 去重比较难实现，我的想法是通过字典加上每个数组hash
func threeSum(nums []int) [][]int {
	l := len(nums)
	res := make([][]int,0)
	for i := 0; i < l; i++ {
		temp := make([]int,3)
		for j := i+1; j < l; j++ {
			for k:=j+1;k<l;k++ {
				if nums[i] + nums[j] +nums[k] == 0{
					temp[0] = nums[i]
					temp[1] = nums[j]
					temp[2] = nums[k]
					res = append(res,temp)
				}
			}
		}
	}
	return res
}

func optimaize_threeSum(nums []int) [][]int {
	res := [][]int{}
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}
	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)
	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}
