package FindDiagonalOrder

func findDiagonalOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return nil
	}

	row, col := len(matrix), len(matrix[0])
	i, j := 0, 0
	res := []int{}

	for {
		// 记录数据
		res = append(res, matrix[i][j])

		// 向右上方走：行标减小，列标增大
		for i-1 >= 0 && j+1 < col {
			i--
			j++
			res = append(res, matrix[i][j])
		}

		// 顺时针转向：要么选右方，要么选下方；
		if j+1 < col {
			j++
		} else if i+1 < row {
			i++
		} else {
			// 无节点可选
			break
		}

		// 记录数据
		res = append(res, matrix[i][j])

		//  向左下方走：行标增大，列标减小
		for i+1 < row && j-1 >= 0 {
			i++
			j--
			res = append(res, matrix[i][j])
		}

		// 逆时针转向：要么选下方，要么右方;
		if i+1 < row {
			i++
		} else if j+1 < col {
			j++
		} else {
			// 无节点可选
			break
		}

	}
	return res

}
