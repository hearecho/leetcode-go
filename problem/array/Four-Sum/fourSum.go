package Four_Sum

import "sort"

/**
⽤ map 提前计算好任意 3 个数字之和，保存起来，可以将时间复杂度降到 O(n^3)。这⼀题⽐较麻烦的
⼀点在于，最后输出解的时候，要求输出不重复的解。数组中同⼀个数字可能出现多次，同⼀个数字也
可能使⽤多次，但是最后输出解的时候，不能重复。例如 [-1，1，2, -2] 和 [2, -1, -2, 1]、[-2, 2, -1, 1] 这
3 个解是重复的，即使 -1, -2 可能出现 100 次，每次使⽤的 -1, -2 的数组下标都是不同的。
 */
func fourSum(nums []int, target int) [][]int {
	//存储结果
	res := [][]int{}
	counter := map[int]int{}
	//数字的重复次数
	for _, value := range nums {
		counter[value]++
	}
	uniqNums := []int{}
	//不重复数字
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)
	for i := 0; i < len(uniqNums); i++ {
		//排除一种情况，四种元素都是一样的的解
		if (uniqNums[i]*4 == target) && counter[uniqNums[i]] >= 4 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i],
				uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*3+uniqNums[j] == target) && counter[uniqNums[i]] > 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i],
					uniqNums[j]})
			}
			if (uniqNums[j]*3+uniqNums[i] == target) && counter[uniqNums[j]] > 2 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j],
					uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i]*2 == target) && counter[uniqNums[j]] > 1 && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j],
					uniqNums[j]})
			}
			for k := j + 1; k < len(uniqNums); k++ {
				if (uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target) &&
					counter[uniqNums[i]] > 1 {
					res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j],
						uniqNums[k]})
				}
				if (uniqNums[j]*2+uniqNums[i]+uniqNums[k] == target) &&
					counter[uniqNums[j]] > 1 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j],
						uniqNums[k]})
				}
				if (uniqNums[k]*2+uniqNums[i]+uniqNums[j] == target) &&
					counter[uniqNums[k]] > 1 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k],
						uniqNums[k]})
				}
				c := target - uniqNums[i] - uniqNums[j] - uniqNums[k]
				if c > uniqNums[k] && counter[c] > 0 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], c})
				}
			}
		}
	}
	return res
}

func optimaize_fourSum(nums []int, target int) [][]int {
	res := make([][]int,0)
	l := len(nums)
	//当数组为nil或者是数组长度小于4则直接返回
	if nums==nil ||  l<4{
		return res
	}
	//排序
	sort.Ints(nums)
	//定义4个指针k，i，j，h  k从0开始遍历，i从k+1开始遍历，留下j和h，j指向i+1，h指向数组最大值
	for k:=0;k<l-3;k++ {
		//去重
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		//获取当前的最小值，如果最小值比目标值大，而排序数组后面的值更大，所以会一直大于目标值
		//上述遍历范围为k-3的原因
		min1 := nums[k]+nums[k+1]+nums[k+2]+nums[k+3]
		if min1 > target {
			break
		}
		//最后三个是数组的最大值，所以如果当前值加上最后三个值还是比target小，则当前值加上后买你随意三个值的和都是小的所以排除
		max1 := nums[k]+nums[l-1]+nums[l-2]+nums[l-3]
		if max1 < target{
			continue
		}
		for i:=k+1;i<l-2;i++ {
			//去重
			if i>k+1 && nums[i] == nums[i+1] {
				continue
			}
			j := i+1
			h := l-1
			min := nums[k] + nums[i] + nums[j] +nums[j+1]
			if min > target {
				break
			}
			max := nums[k] + nums[i] + nums[h]+ nums[h-1]
			if max < target {
				continue
			}
			//双指针聚合
			for j < h{
				curr := nums[k]+nums[i]+nums[j]+nums[h]
				if curr == target {
					temp := []int{nums[k],nums[i],nums[j],nums[h]}
					res = append(res,temp)
					j++
					//去重
					for j<h&& nums[j] == nums[j-1] {
						j++
					}
					h--
					//去重
					for j<h && i<h && nums[h]==nums[h+1] {
						h--
					}
				} else if curr > target {
					h--
				} else {
					j++
				}
			}
		}

	}
	return res
}
