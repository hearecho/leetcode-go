package Three_Sum_Closest

import (
	"math"
	"sort"
)

//类似三数之和的暴力解法
func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	res, diff := 0, math.MaxInt64
	//排序 从小到大
	sort.Ints(nums)
	//遍历
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				temp := nums[i] + nums[j] + nums[k] - target
				if abs(temp) < diff {
					diff = abs(temp)
					res = temp + target
				}
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//优化解法
//优化方法，确定第一个数之后，然后使用双指针进行二分查找
func optimaize_threeSumClosest(nums []int, target int) int {
	l := len(nums)
	//排序
	sort.Ints(nums)
	//变量
	res, diff := 0, math.MaxInt64
	for i := 0; i < l; i++ {
		//二分查找
		for j, k := i+1, l-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			//首先要记录最小的差值，之后才能进行二分
			if abs(sum-target) < diff {
				res, diff = sum, abs(sum-target)
			}
			//二分
			if sum == target {
				return res
			} else if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
