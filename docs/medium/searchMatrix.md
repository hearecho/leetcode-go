# [搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/)
> 编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值。该矩阵具有如下特性：

### 注意点
> 每行中的整数从左到右按升序排列。  每行的第一个整数大于前一行的最后一个整数。

### 解题思路
> 由于数组从左到右升序，从上到下升序,所以可以从右上角开始遍历

```go
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
	i,j := 0,len(matrix[0])-1
	for i < len(matrix) && j >=0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
```