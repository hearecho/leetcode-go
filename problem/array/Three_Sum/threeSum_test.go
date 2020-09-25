package Three_Sum

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
}
func TestThreeSum(t *testing.T)  {
	params := []param{
		{[]int{-1, 0, 1, 2, -1, -4}},
	}
	for _,p := range params {
		fmt.Printf("【input】:%v       【output】:%v\n",p,optimaize_threeSum(p.nums))
	}
}
