package GetPermutation

import (
	"strconv"
)

func getPermutation(n int, k int) string {
	res := ""
	used := make([]bool, n)
	index := 0
	backtrack("", &res, n, &used, &index, k)
	return res
}

func backtrack(curStr string, res *string, n int, used *[]bool, index *int, k int) {
	if *index > k {
		return
	}
	if len(curStr) == n {
		if *index == k-1 {
			*res = curStr
		}
		*index++
		return
	}
	for i := 1; i <= n; i++ {
		if !(*used)[i-1] {
			(*used)[i-1] = true
			curStr = curStr + strconv.Itoa(i)
			backtrack(curStr, res, n, used, index, k)
			(*used)[i-1] = false
			curStr = curStr[:len(curStr)-1]
		}
	}
}
