package SpiralOrder

import (
	"fmt"
	"testing"
)

type param struct {
	matrix [][]int
}

func Test_spiralOrder(t *testing.T) {
	params := []param{
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, spiralOrder(p.matrix))
	}
}
