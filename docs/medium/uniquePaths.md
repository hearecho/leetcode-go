# 不同路径
> 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
>机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
>问总共有多少条不同的路径？
>
### 解题思路
![Bg38jH.png](https://s1.ax1x.com/2020/11/04/Bg38jH.png)
```go
func uniquePaths(m int, n int) int {
	dp := make([][]int,m)
	for i:=0;i<m;i++ {
		dp[i] = make([]int,n)
	}
	//init
	for i:=0;i<m;i++ {
		dp[i][0] = 1
	}
	for i:=0;i<n;i++ {
		dp[0][i] = 1
	}
	for i := 1;i<m;i++ {
		for j:=1;j<n;j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
```