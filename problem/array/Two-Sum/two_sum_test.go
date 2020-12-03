package Two_Sum

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
	target int
}
func Test_Problem(t *testing.T)  {
	params := []param{
		{[]int{3, 2, 4}, 6},
		{[]int{3, 2, 4}, 5},
		{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
		{[]int{0, 1}, 1},
		{[]int{0, 3}, 5},
		{[]int{3,3},6},
	}
	for _,p := range params {
		fmt.Printf("【input】:%v       【output】:%v\n",p,twoSum_optimize(p.nums,p.target))
	}
}
