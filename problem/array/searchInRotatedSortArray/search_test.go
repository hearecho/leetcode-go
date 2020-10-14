package searchInRotatedSortArray

import (
	"fmt"
	"testing"
)

type param struct {
	nums   []int
	target int
}

func Test_Search(t *testing.T) {
	params := []param{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
	}
	for _, p := range params {
		fmt.Printf("【input】：%v\t【output】：%v\n", p, search(p.nums, p.target))
	}
}
