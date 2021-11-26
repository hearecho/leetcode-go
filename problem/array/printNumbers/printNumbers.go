package printNumbers

import (
	"strconv"
)

func printNumbers(n int) []int {
	// 可以简单理解为全排列问题
	// 例如 n为2
	// 则 十位可以取值 [0~9]  个位取值 [0~9]
	res := []int{}
	var bk func(s string)
	bk = func(s string) {
		if len(s) == n {
			temp, _ := strconv.Atoi(s)
			res = append(res, temp)
			return
		}
		for i := 0; i <= 9; i++ {
			bk(s + strconv.Itoa(i))
		}
	}
	bk("")
	return res[1:]
}
