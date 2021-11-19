package MovingCount

func movingCount(m int, n int, k int) int {
	// 回溯搜索
	// 同时还需要记录这个位置是否已经被走过
	res := 0
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	moves := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	var bk func(i, j int)
	bk = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || k < calSum(i, j) || visited[i][j] {
			return
		}
		visited[i][j] = true
		res++
		for _, move := range moves {
			bk(i+move[0], j+move[1])
		}
	}
	bk(0, 0)
	return res
}

func calSum(i, j int) int {
	// 计算坐标的位数之和
	sum := 0
	for i > 0 {
		sum += i % 10
		i = i / 10
	}
	for j > 0 {
		sum += j % 10
		j = j / 10
	}
	return sum
}
