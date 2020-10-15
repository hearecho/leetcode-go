package findfirstlastIndex

import (
	"fmt"
	"testing"
)

type param struct {
	nums   []int
	target int
}

func Test_searchRange(t *testing.T) {
	params := []param{
		{[]int{5, 7, 7, 8, 8, 10}, 8},
		{[]int{5, 7, 7, 8, 8, 10}, 6},
		{[]int{2, 2}, 2},
		{[]int{1}, 1},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, searchRange(p.nums, p.target))
	}
}
