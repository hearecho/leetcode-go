package dailyProblem

//https://leetcode-cn.com/problems/candy/
//分糖果问题 ，贪心
//注意点 1. 每个孩子至少一个糖果 2. 相邻孩子必须获得更多的糖果
func candy(ratings []int) int {
	//
	left := make([]int, len(ratings))
	res := 0
	//先遍历一遍排序,如果 左边的人的分数小于右边的分数，则左边的糖果定位1
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	//之后再遍历一次，从后往前遍历，
	right := 0
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		res += max(left[i], right)
	}
	return res
}
