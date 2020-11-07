package dailyProblem
/*
给定一个整数数组nums，返回区间和在[lower, upper]之间的个数，包含lower和upper。
区间和S(i, j)表示在nums中，位置从i到j的元素之和，包含i和j(i ≤ j)。
 */
//暴力法 O(n2)
func countRangeSum(nums []int, lower int, upper int) int {
	count := 0

	for i:=0;i<len(nums);i++ {
		res := nums[i]
		for j:=i;j<len(nums);j++ {
			if j == i {
				res = nums[i]
			} else {
				res += nums[j]
			}
			if res >= lower && res <= upper {
				count++
			}
		}
	}
	return count
}
