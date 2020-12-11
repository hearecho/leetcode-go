package dailyProblem

import "fmt"

//https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/
func isPossible(nums []int) bool {
	//数组是按照升序进行排序的
	//由于要求是子序列长度至少为3，当我们的子序列大于3的时候，后续增加的数字个数不能为1，如果小于1则停止增加
	counter := make([]int, nums[len(nums)-1])
	for _, num := range nums {
		counter[num-1]++
	}
	//设置一个参数，记录counter中的最大值
	max_counter := -1
	for max_counter != 0 {
		temp := make([]int, 0)
		temp_max := -1
		for _, num := range nums {
			if counter[num-1] > 0 {
				if len(temp) != 0 && num == temp[len(temp)-1] {
					continue
				}
				if len(temp) >= 3 && counter[num-1] <= 1 {
					break
				}
				temp = append(temp, num)
				counter[num-1]--
				if counter[num-1] > temp_max {
					temp_max = counter[num-1]
				}
			}
		}
		max_counter = temp_max
		fmt.Println(temp)
		if len(temp) < 3 {
			return false
		}
	}
	return true
}
