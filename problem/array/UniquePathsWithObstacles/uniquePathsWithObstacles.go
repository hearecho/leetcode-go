package UniquePathsWithObstacles

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	//原地修改
	for i := 0; i < len(obstacleGrid); i++ {
		//边界上有阻碍，则后面的都设置为0
		if obstacleGrid[i][0] != 1  {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
			for j:=i+1;j<len(obstacleGrid);j++ {
				obstacleGrid[j][0] = 0
			}
			break
		}
	}
	for i := 1; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] != 1 {
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
			for j:=i+1;j<len(obstacleGrid[0]);j++ {
				obstacleGrid[0][j] = 0
			}
			break
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] != 1 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
