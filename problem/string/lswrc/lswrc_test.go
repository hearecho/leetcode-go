package lswrc

import (
	"fmt"
	"testing"
)

type param struct {
	s string
}
func Test_lengthOfLongestSubstring(t *testing.T)  {
	params := []param{
		{"abcabcbb"},
		{"bbbbb"},
		{"pwwkew"},
		{""},
	}

	for _,p := range params {
		lengthOfLongestSubstring(p.s)
		fmt.Printf("【input】:%v       【output】:%v\n",p,lengthOfLongestSubstring(p.s))
	}
}
