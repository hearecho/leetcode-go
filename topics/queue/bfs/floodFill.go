package bfs

//var moves = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	r, c := len(image), len(image[0])
	color := image[sr][sc]
	if color == newColor {
		return image
	}
	queue := make([][]int, 0)
	queue = append(queue, []int{sr, sc})
	for len(queue) != 0 {
		//四个方向
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[i]
			image[point[0]][point[1]] = newColor
			for j := 0; j < len(moves); j++ {
				nr, nc := point[0]+moves[j][0], point[1]+moves[j][1]
				if nr >= 0 && nr < r && nc >= 0 && nc < c && image[nr][nc] == color {
					queue = append(queue, []int{nr, nc})
				}
			}
		}
		queue = queue[size:]
	}
	return image
}
