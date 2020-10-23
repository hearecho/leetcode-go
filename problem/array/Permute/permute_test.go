package Permute

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
}

func Test_permute(t *testing.T) {
	params := []param{
		{[]int{1, 2, 3}},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【ouput】:%v\n", p, permute(p.nums))
	}
}
