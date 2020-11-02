package GetPermutation

import (
	"fmt"
	"testing"
)

type param struct {
	n, k int
}

func Test_getPermutation(t *testing.T) {
	params := []param{
		{3, 3},
		{4, 9},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, getPermutation(p.n, p.k))
	}
}
