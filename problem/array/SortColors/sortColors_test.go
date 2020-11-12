package SortColors

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
}

func Test_sortColors(t *testing.T)  {
	params := []param{
		{[]int{2,0,2,1,1,0}},
	}
	for _,p := range params{
		temp := p.nums
		sortColors(p.nums)
		fmt.Printf("【input】:%v\t【output】:%v\n",temp,p.nums)
	}
}