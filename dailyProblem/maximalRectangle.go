package dailyProblem

func maximalRectangle(m [][]byte) int {
	//动态规划
	res := 0
	rows, cols := len(m), len(m[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	//第一行 第一列单独讨论
	for i := 0; i < rows; i++ {
		//第一列
		if i-1 > 0 {
			if m[i][0] == 1 {
				//dp[i][0] =
			}
		} else {

		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			//以(i,j)为右下端点形成的矩形的面积，只有为1的时候才进行判断
			if m[i][j] == 1 {

			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return res
}
