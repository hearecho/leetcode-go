package Permute2

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
}

func Test_permuteUnique(t *testing.T) {
	params := []param{
		{[]int{1, 2, 1}},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【ouput】:%v\n", p, permuteUnique(p.nums))
	}
}
