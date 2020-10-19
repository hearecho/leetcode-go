package solveshudu

func solveSudoku(board [][]byte) {
	//行列中标识数据是否已经存在
	var line, column [9][9]bool
	//标识一个3*3块中元素存在数据
	var block [3][3][9]bool
	//标识空格
	var spaces [][2]int

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				//增加空格元素
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				///标识行
				line[i][digit] = true
				//标识列第二个元素是存储的数据
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(int) bool
	//递归回溯
	dfs = func(pos int) bool {
		//空格遍历完成遍历结束
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		//回溯主体
		for digit := byte(0); digit < 9; digit++ {
			//如果这个元素在行，列，以及3*3方格中均不存在，那么便可以进一步的回溯
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				//填入之后设置相应的数字为已存
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1'
				if dfs(pos + 1) {
					return true
				}
				//回溯就是需要对于改变过的状态再将其改变回来，如果没有改变状态可以不管
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	//遍历空白位置集合，从第0个开始
	dfs(0)
}
