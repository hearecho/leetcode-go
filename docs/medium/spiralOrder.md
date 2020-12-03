# 螺旋矩阵
> 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

### 解题思路
> 我们可以一层一层的遍历，直到遍历到内层结束。
> 所以需要一个结束条件，因为矩阵不是方阵
> 并且需要控制从遍历方向的转变

![](https://i.ibb.co/fdCmttf/04b87d378198d701c4e2572e07259d3.jpg)
```go
func spiralOrder(matrix [][]int) []int {
	tempRes := make([]int, 0)
	if len(matrix) == 0 {
		return tempRes
	}
	r1, r2, c1, c2 := 0, len(matrix)-1, 0, len(matrix[0])-1
	for r1 <= r2 && c1 <= c2 {
		//遍历上行，从左往右
		for i := c1; i <= c2; i++ {
			tempRes = append(tempRes, matrix[r1][i])
		}
		//遍历右列，从第二行开始,从上往下
		for i := r1 + 1; i <= r2; i++ {
			tempRes = append(tempRes, matrix[i][c2])
		}
		//遍历下行
		if r2 > r1 {
			for i := c2 - 1; i >= c1; i-- {
				tempRes = append(tempRes, matrix[r2][i])
			}
		}
		//遍历左列从下往上
		if c2 > c1 {
			for i := r2 - 1; i > r1; i-- {
				tempRes = append(tempRes, matrix[i][c1])
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