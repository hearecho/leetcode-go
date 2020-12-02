package FindDiagonalOrder

import (
	"fmt"
	"testing"
)

type param struct {
	matrix [][]int
}

func Test_findDiagonalOrder(t *testing.T) {
	params := []param{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, findDiagonalOrder(p.matrix))
	}
}
