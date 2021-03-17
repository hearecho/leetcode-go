package dailyProblem

//动态规划
//		b	a	b	g	b	a	g
//	b	1	0	1	0	1	0	0
//	a	0	1	0	0	0	3	0
//	g	0	0	0	1	0	0	4
// 遍历最后一行将数值相加等到 5

func numDistinct(s string, t string) int {
	l_s, l_t := len(s), len(t)
	dp := make([][]int, l_t)
	for i := 0; i < l_t; i++ {
		dp[i] = make([]int, l_s)
	}
	//初始化第一行
	for i := 0; i < l_s; i++ {
		if s[i] == t[0] {
			dp[0][i] = 1
		}
	}
	//动态规划数组创建
	//只有当此字母和原串字母相同时，才向里面进行查找，并相加
	for i := 1; i < l_t; i++ {
		ch := t[i]
		for j := 0; j < l_s; j++ {
			if ch == s[j] {
				for k := 0; k < j; k++ {
					dp[i][j] += dp[i-1][k]
				}
			}
		}
	}
	//处理结果
	res := 0
	for i := 0; i < l_s; i++ {
		res += dp[l_t-1][i]
	}
	return res
}
