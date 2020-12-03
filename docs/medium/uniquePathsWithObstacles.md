# 不同路径 II
> 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

### 注意点
> 边界处理，边界需要判断如果有阻塞则后面都是不可以到达的，所以需要注意
> 可以原地修改数据,而不使用额外数组

### 解题思路
```go
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

```