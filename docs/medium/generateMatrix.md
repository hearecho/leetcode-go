# 螺旋矩阵2
> 给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

### 解题思路
> 这道题是螺旋矩阵的逆方法，这个是构建数组，那我们只需要按照顺序直接模拟便可以了

```go
func generateMatrix(n int) [][]int {
	tempRes := make([][]int, n)
	for i := 0; i < n; i++ {
		tempRes[i] = make([]int, n)
	}

	r1, r2, c1, c2 := 0, n-1, 0, n-1
	k := 1
	for r1 <= r2 && c1 <= c2 {
		//上行
		for i := c1; i <= c2; i++ {
			tempRes[r1][i] = k
			k++
		}
		//右列
		for i := r1 + 1; i <= r2; i++ {
			tempRes[i][c2] = k
			k++
		}
		//下行左列
		if r1 < r2 && c1 < c2 {
			for i := c2 - 1; i >= c1; i-- {
				tempRes[r2][i] = k
				k++
			}
			for i := r2 - 1; i > r1; i-- {
				tempRes[i][c1] = k
				k++
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return tempRes
}

```