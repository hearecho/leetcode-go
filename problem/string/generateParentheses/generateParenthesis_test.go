package generateParentheses

import (
	"fmt"
	"testing"
)

type param struct {
	n int
}

func Test_generateParenthesis(t *testing.T)  {
	params := []param{
		{0},
		{1},
		{2},
		{3},
		{4},
	}
	for _,p := range params{
		fmt.Printf("【input】:%v       【output】:%v\n",p,generateParenthesis(p.n))
	}
}
