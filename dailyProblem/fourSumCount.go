package dailyProblem

//https://leetcode-cn.com/problems/4sum-ii/
//数组具有相同的长度
//回溯  超出时间复杂度
func fourSumCount(A []int, B []int, C []int, D []int) int {
	g := make([][]int, 4)
	g[0], g[1], g[2], g[3] = A, B, C, D
	res := 0
	backtrack(g, 0, 0, &res)
	return res
}

func backtrack(g [][]int, usedNum int, cur int, res *int) {
	//回溯结束
	if usedNum == 4 {
		if cur == 0 {
			*res++
		}
		return
	}
	for i := 0; i < len(g[usedNum]); i++ {
		backtrack(g, usedNum+1, cur+g[usedNum][i], res)
	}
}

//哈希表 分组先存储两个数的和，之后再后面找剩下两个数的和相加为0即可
func fourSumCount1(a, b, c, d []int) (ans int) {
	countAB := map[int]int{}
	for _, v := range a {
		for _, w := range b {
			countAB[v+w]++
		}
	}
	for _, v := range c {
		for _, w := range d {
			ans += countAB[-v-w]
		}
	}
	return
}
