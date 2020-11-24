package bfs

var moves = [][]int{
	{-1, 0}, {1, 0}, {0, 1}, {0, -1},
}

func NumIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				bfs(i, j, grid)
				res++
			}
		}
	}
	return res
}

func bfs(x, y int, grid [][]byte) {
	queue := make([][]int, 0)
	queue = append(queue, []int{x, y})
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[0]
			grid[point[0]][point[1]] = 0
			for _, move := range moves {
				nx, ny := move[0]+point[0], move[1]+point[1]
				if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) || grid[nx][ny] == 0 {
					continue
				}
				queue = append(queue, []int{nx, ny})
			}

			queue = queue[1:]
		}
	}
}
