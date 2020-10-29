package CanJump

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
}

func Test_canJump(t *testing.T) {
	params := []param{
		{[]int{2, 0, 1, 1, 4}},
		{[]int{3, 2, 1, 0, 4}},
		{[]int{0}},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【ouput】:%v\n", p, canJump(p.nums))
	}
}
