# 旋转图像
> 给定一个 n × n 的二维矩阵表示一个图像。
> 将图像顺时针旋转 90 度。
> 说明：你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

### 注意点
1. 这个二维矩阵是方阵

### 解题思路
> 由于矩阵是方针，所以我们可以先将方针上下呼唤，之后在对角进行交换便可以完成顺时针旋转

![](http://riyugo.com/i/2020/10/24/ou9ude.jpg)

```go
func rotate(matrix [][]int) {
	//上下互换
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			swap(&matrix[i][j], &matrix[len(matrix)-i-1][j])
		}
	}
	//对角线对称互换
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			swap(&matrix[i][j], &matrix[j][i])
		}
	}
}

func swap(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}
```
