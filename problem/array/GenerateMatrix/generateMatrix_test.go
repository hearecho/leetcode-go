package GenerateMatrix

import (
	"fmt"
	"testing"
)

type param struct {
	n int
}

func Test_generateMatrix(t *testing.T) {
	params := []param{
		{3},
		{4},
		{5},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, generateMatrix(p.n))
	}
}
