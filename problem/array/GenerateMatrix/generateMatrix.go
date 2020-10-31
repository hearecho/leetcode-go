package GenerateMatrix

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	r1, r2, c1, c2 := 0, n-1, 0, n-1
	k := 1
	for r1 <= r2 && c1 <= c2 {
		//上行
		for i := c1; i <= c2; i++ {
			res[r1][i] = k
			k++
		}
		//右列
		for i := r1 + 1; i <= r2; i++ {
			res[i][c2] = k
			k++
		}
		//下行左列
		if r1 < r2 && c1 < c2 {
			for i := c2 - 1; i >= c1; i-- {
				res[r2][i] = k
				k++
			}
			for i := r2 - 1; i > r1; i-- {
				res[i][c1] = k
				k++
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return res
}
