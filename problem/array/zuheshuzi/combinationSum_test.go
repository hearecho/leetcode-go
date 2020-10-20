package zuheshuzi

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
		{[]int{2, 3, 6, 7}, 7},
		{[]int{2, 3, 5}, 8},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, combinationSum(p.nums, p.target))
	}
}
