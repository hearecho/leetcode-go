package ImplementStrstr

import (
	"fmt"
	"testing"
)

type param struct {
	haystack string
	needle string
}

func Test_Strstr(t *testing.T)  {
	params := []param{
		{"hello","ll"},
		{ "aaaaa", "bba"},
	}
	for _,p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n",p,strStr_2(p.haystack,p.needle))
	}
}