# 最小路径和
> 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
>
### 注意点
> 每次只能向下或者是向右移动一步
>
### 解题思路
> 典型的动态规划题目，优化可以使用现有的数组
>递推公式 `dp[i] =min(dp[i-1][j],dp[i][j-1])+dp[i]` 
>
```go
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
```