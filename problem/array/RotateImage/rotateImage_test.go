package RotateImage

import (
	"fmt"
	"testing"
)

type param struct {
	matrix [][]int
}

func Test_rotate(t *testing.T) {
	params := []param{
		{[][]int{{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9}}},
	}
	for _, p := range params {
		temp := p.matrix
		rotate(p.matrix)
		fmt.Printf("【input】:%v\t【ouput】:%v\n", temp, p.matrix)
	}
}
func Test_swap(t *testing.T) {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
