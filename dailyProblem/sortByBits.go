package dailyProblem

import "sort"

/*
给你一个整数数组arr。请你将数组中的元素按照其二进制表示中数字 1 的数目升序排序。
如果存在多个数字二进制中1的数目相同，则必须将它们按照数值大小升序排列。
请你返回排序后的数组。
 */
func sortByBits(arr []int) []int {
	count := make(map[int]int,len(arr))
	for i:=0;i<len(arr);i++ {
		count[arr[i]] = count_one_num(arr[i])
	}
	sort.Slice(arr, func(i, j int) bool {
		if count[arr[i]] < count[arr[j]] {
			return true
		}
		if count[arr[i]] == count[arr[j]] && arr[i] < arr[j]{
			return true
		}
		return false
	})
	return arr
}

/*
返回一个数的二进制中1的个数
 */
func count_one_num(x int) int  {
	count := 0
	for x > 0 {
		x &= x-1
		count++
	}
	return count
}
