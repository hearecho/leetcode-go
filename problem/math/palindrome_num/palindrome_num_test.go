package palindrome_num

import (
	"fmt"
	"testing"
)

type param struct {
	x int
}

func Test_isPalindrome(t *testing.T)  {
	ps := []param {
		{121},
		{-121},
		{0},
		{789},
		{10},
		{21120},
	}
	for _,p := range ps {
		fmt.Printf("【input】:%v       【output】:%v\n",p,optimaize_isPalindrome(p.x))
	}
}
