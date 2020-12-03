package LongestPalindrome

import (
	"fmt"
	"testing"
)

type param struct {
	s string
}

func Test_longestPalindrome(t *testing.T) {
	params := []param{
		{"abbabb"},
		{"aba"},
		{"aaaa"},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, longestPalindrome(p.s))
	}
}
