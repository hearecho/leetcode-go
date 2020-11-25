# 第k个排列
> 给出集合 `[1,2,3,…,n]`，其所有元素共有 n! 种排列。
> 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：`"123""132""213""231""312""321"`
> 给定 n 和 k，返回第 k 个排列。

### 解题思路
> 全排列问题，但是需要注意时间超出限制，所以需要进行剪枝，不采用数组存储产生的字符串，而是在记录产出字符串的个数，在之后的数字都直接剪去

```go
func getPermutation(n int, k int) string {
	tempRes := ""
	used := make([]bool, n)
	index := 0
	backtrack("", &tempRes, n, &used, &index, k)
	return tempRes
}

func backtrack(curStr string, tempRes *string, n int, used *[]bool, index *int, k int) {
	if *index >= k {
		return
	}
	if len(curStr) == n {
		if *index == k-1 {
			*tempRes = curStr
		}
		*index++
		return
	}
	for i := 1; i <= n; i++ {
		if !(*used)[i-1] {
			(*used)[i-1] = true
			curStr = curStr + strconv.Itoa(i)
			backtrack(curStr, tempRes, n, used, index, k)
			(*used)[i-1] = false
			curStr = curStr[:len(curStr)-1]
		}
	}
}
```