package SortedArrayToBST

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
}

func Test_sortedArrayToBST(t *testing.T)  {
	params := []param{
		{[]int{-10,-3,0,5,9}},
	}
	for _,p := range params {
		root := sortedArrayToBST(p.nums)
		fmt.Printf("【input】:%v\t【output】:%v\n",p.nums,root.Val)
	}
}