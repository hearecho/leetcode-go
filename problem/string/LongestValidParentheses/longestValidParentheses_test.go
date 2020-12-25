package LongestValidParentheses

import (
	"fmt"
	"testing"
)

type param struct {
	s string
}

func Test_longestValidParentheses(t *testing.T) {
	params := []param{
		{"())"},
		{")()())"},
		{"((())())()"},
		{"()(()"},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, longestValidParentheses(p.s))
	}
}
