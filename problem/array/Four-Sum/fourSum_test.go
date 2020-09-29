package Four_Sum

import (
	"fmt"
	"testing"
)

type param struct {
	nums   []int
	target int
}

func TestFourSum(t *testing.T)  {
	params := []param{
		{[]int{1, 0, -1, 0, -2, 2},0},
	}
	for _,p := range params{
		fmt.Printf("【input】:%v       【output】:%v\n",p,optimaize_fourSum(p.nums,p.target))
	}
}
