package SearchInsertPostion

import (
	"fmt"
	"testing"
)

type param struct {
	nums   []int
	target int
}

func Test_searchInsert(t *testing.T) {
	params := []param{
		{[]int{1, 3, 5, 6}, 5},
		{[]int{1, 3, 5, 6}, 2},
		{[]int{1, 3, 5, 6}, 7},
		{[]int{1, 3, 5, 6}, 0},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, searchInsert(p.nums, p.target))
	}
}
