package MinPathSum

func minPathSum(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])
	for i:=0;i<r;i++ {
		if i == 0 {
			continue
		}
		grid[i][0] += grid[i-1][0]
	}
	for i:=0;i<c;i++ {
		if i == 0 {
			continue
		}
		grid[0][i] += grid[0][i-1]
	}
	for i:=1;i<r;i++ {
		for j:=1;j<c;j++ {
			grid[i][j] += min(grid[i-1][j],grid[i][j-1])
		}
	}
	return grid[r-1][c-1]
}

func min(a,b int) int  {
	if a > b{
		return b
	}
	return  a
}