package shudu

import "strconv"

/**
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
数独部分空格内已填入了数字，空白格用 '.' 表示。
 */
func isValidSudoku(board [][]byte) bool {
	//按照中心先判断九个方格是否符合
	for i :=1;i<9;{
		for j:=1;j<9; {
			if !checkBox(i,j,board) {
				return false
			}
			j+=3
		}
		i+=3
	}
	//判断行，列
	for i:=0;i<9;i++ {
		counter := make([]int,9)
		for j:=0;j<9;j++ {
			if board[i][j] == '.' {
				continue
			}
			num,_ := strconv.Atoi(string(board[i][j]))
			counter[num-1]++
			if counter[num-1] >= 2 {
				return false
			}
		}
	}
	//判断列
	for j:=0;j<9;j++ {
		counter := make([]int,9)
		for i:=0;i<9;i++ {
			if board[i][j] == '.' {
				continue
			}
			num,_ := strconv.Atoi(string(board[i][j]))
			counter[num-1]++
			if counter[num-1] >= 2 {
				return false
			}
		}
	}
	return true
}

func checkBox(a,b int,board [][]byte) bool {
	counter := make([]int,9)
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++ {
			if board[a-1+i][b-1+j] == '.' {
				continue
			}
			num,_ := strconv.Atoi(string(board[a-1+i][b-1+j]))
			counter[num-1]++
			if counter[num-1] >= 2 {
				return false
			}
		}
	}
	return true
}
