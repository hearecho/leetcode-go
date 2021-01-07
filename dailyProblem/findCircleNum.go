package dailyProblem

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	res := 0
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			res++
			// 做一个广度优先搜索 之后将搜索到的标
			queue := make([]int, 0)
			queue = append(queue, i)
			visited[i] = true
			for len(queue) != 0 {
				node := queue[0]
				for j := 0; j < len(isConnected); j++ {
					if isConnected[node][j] == 1 && node != j && !visited[j] {
						queue = append(queue, j)
						visited[j] = true
					}
				}
				queue = queue[1:]
			}
		}
	}
	return res
}
