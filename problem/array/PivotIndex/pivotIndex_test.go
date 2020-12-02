package PivotIndex

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
}

func Test_pivotIndex(t *testing.T) {
	params := []param{
		{[]int{1, 2, 3}},
		{[]int{1, 2, 1}},
		{[]int{1, 7, 3, 6, 5, 6}},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, pivotIndex(p.nums))
	}
}
