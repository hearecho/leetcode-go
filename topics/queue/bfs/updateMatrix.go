package bfs

func UpdateMatrix(matrix [][]int) [][]int {
	var bfs func(r, c int) int
	res := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		res[i] = make([]int, len(matrix[0]))
	}
	bfs = func(sr, sc int) int {
		r, c := len(matrix), len(matrix[0])
		res := 0
		queue := make([][]int, 0)
		queue = append(queue, []int{sr, sc})
		for len(queue) != 0 {
			size := len(queue)
			res++
			for i := 0; i < size; i++ {
				point := queue[i]
				for j := 0; j < len(moves); j++ {
					nr, nc := point[0]+moves[j][0], point[1]+moves[j][1]
					if nr >= 0 && nr < r && nc >= 0 && nc < c {
						if matrix[nr][nc] == 0 {
							return res
						}
						queue = append(queue, []int{nr, nc})
					}
				}
			}
			queue = queue[size:]
		}
		return -1
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			step := 0
			if matrix[i][j] != 0 {
				step = bfs(i, j)
			}
			res[i][j] = step
		}
	}
	return res
}
