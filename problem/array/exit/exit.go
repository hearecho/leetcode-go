package exit

func exist(board [][]byte, word string) bool {
	// 在搜索过程中与word进行比较
	// 搜索过程是一个不断回溯的过程
	moves := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	res := false
	// 直接将原数组字符修改为空字符
	var bk func(curLen, i, j int)
	bk = func(curLen, i, j int) {
		// 停止条件
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[curLen] {
			return
		}
		if curLen == len(word)-1 {
			res = res || true
			return
		}
		for _, move := range moves {
			board[i][j] = ' '
			bk(curLen+1, i+move[0], j+move[1])
			board[i][j] = word[curLen]
		}

	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			bk(0, i, j)
			if res {
				return res
			}
		}
	}
	return res
}
