package dailyProblem

type NumMatrix struct {
	numMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{numMatrix: matrix}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := 0
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			res += m.numMatrix[i][j]
		}
	}
	return res
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
