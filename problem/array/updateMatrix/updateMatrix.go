package updateMatrix

type Loc struct {
	X int
	Y int
}

func updateMatrix(mat [][]int) [][]int {
	// 最近距离  一般求最近距离 使用广度优先搜索
	// 如果当前值为0 则直接赋值0
	res := make([][]int, len(mat))
	for i := range res {
		res[i] = make([]int, len(mat[0]))
	}
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var bfs func(x, y int)
	bfs = func(x, y int) {
		visited := make([][]bool, len(mat))
		for i := range visited {
			visited[i] = make([]bool, len(mat[0]))
		}
		distance := 0
		queue := make([]Loc, 0)
		queue = append(queue, Loc{x, y})
		visited[x][y] = true
		for len(queue) > 0 {
			l := len(queue)
			distance++
			for i := 0; i < l; i++ {
				loc := queue[i]
				for _, move := range moves {
					nx, ny := loc.X+move[0], loc.Y+move[1]
					if nx >= 0 && nx < len(mat) && ny >= 0 && ny < len(mat[0]) && !visited[nx][ny] {
						if mat[nx][ny] == 0 {
							res[x][y] = distance
							return
						}
						queue = append(queue, Loc{nx, ny})
						visited[nx][ny] = true
					}
				}
			}
			queue = queue[l:]
		}
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] != 0 {
				bfs(i, j)
			}
		}
	}
	return res
}

//  直接使用广度优先搜索回超出时间限制
// 我们可以使用一个超级源点的机制
// 并且不是从1进行搜索，而是从0开始搜索1 搜索到1则增加距离
func updateMatrixDP(mat [][]int) [][]int {
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(mat), len(mat[0])
	res := make([][]int, m)
	visited := make([][]bool, m)
	for i := range res {
		res[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}
	queue := make([]Loc, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, Loc{i, j})
				visited[i][j] = true
			}
		}
	}
	for len(queue) > 0 {
		loc := queue[0]
		i, j := loc.X, loc.Y
		for _, move := range moves {
			nx, ny := i+move[0], j+move[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !visited[nx][ny] {
				res[nx][ny] = res[i][j] + 1
				queue = append(queue, Loc{nx, ny})
				visited[nx][ny] = true
			}
		}
		queue = queue[1:]
	}
	return res
}
