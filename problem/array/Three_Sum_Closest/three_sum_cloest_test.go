package Three_Sum_Closest

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
	target int
}
func TestThreeSumCloest(t *testing.T)  {
	params := []param{
		{[]int{-1,2,1,-4},2},
	}

	for _,p := range params {
		fmt.Printf("【input】:%v    【output】:%v",p,optimaize_threeSumClosest(p.nums,p.target))
	}
}
