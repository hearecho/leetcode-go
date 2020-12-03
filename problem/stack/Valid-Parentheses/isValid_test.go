package Valid_Parentheses

import (
	"fmt"
	"testing"
)

type param struct {
	s string
}

func Test_isValid(t *testing.T)  {
	params := []param {
		{"()"},
		{"()[]{}"},
		{"(]"},
		{"([)]"},
		{"{[]}"},
		{"["},
	}

	for _,p := range params {
		fmt.Printf("【input】:%v       【output】:%v\n",p,isValid(p.s))
	}
}