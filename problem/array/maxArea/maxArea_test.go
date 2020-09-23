package maxArea

import (
	"fmt"
	"testing"
)

type param struct {
	height []int
}
func Test_MaxArea(t *testing.T)  {
	ps := []param{
		{[]int{1,8,6,2,5,4,8,3,7}},
		{[]int{2,3,4,5,18,17,6}},
	}

	for _,p := range ps{
		fmt.Printf("【input】:%v       【output】:%v\n",p,maxArea(p.height))
	}
}
