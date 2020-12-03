package Trap

import (
	"fmt"
	"testing"
)

type param struct {
	height []int
}

func Test_trap(t *testing.T) {
	params := []param{
		{[]int{4, 2, 0, 3, 2, 5}},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, trap(p.height))
	}
}
