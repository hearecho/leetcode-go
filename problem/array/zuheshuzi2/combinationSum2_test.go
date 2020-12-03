package zuheshuzi2

import (
	"fmt"
	"testing"
)

type param struct {
	nums   []int
	target int
}

func Test_combinationSum(t *testing.T) {
	params := []param{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
		{[]int{2, 5, 2, 1, 2}, 5},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, combinationSum2(p.nums, p.target))
	}
}
