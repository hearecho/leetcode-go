package SpiralOrder

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}
	r1, r2, c1, c2 := 0, len(matrix)-1, 0, len(matrix[0])-1
	for r1 <= r2 && c1 <= c2 {
		//遍历上行，从左往右
		for i := c1; i <= c2; i++ {
			res = append(res, matrix[r1][i])
		}
		//遍历右列，从第二行开始,从上往下
		for i := r1 + 1; i <= r2; i++ {
			res = append(res, matrix[i][c2])
		}
		//遍历下行
		if r2 > r1 {
			for i := c2 - 1; i >= c1; i-- {
				res = append(res, matrix[r2][i])
			}
		}
		//遍历左列从下往上
		if c2 > c1 {
			for i := r2 - 1; i > r1; i-- {
				res = append(res, matrix[i][c1])
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return res
}
