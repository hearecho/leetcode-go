package letterCombinations

import (
	"fmt"
	"testing"
)

type param struct {
	s string
}
func Test_letterCombinations(t *testing.T)  {
	params := []param{
		//{"456"},
		{"23"},
		{""},
	}

	for _,p := range params {
		fmt.Printf("【input】:%v       【output】:%v\n",p,letterCombinations(p.s))
	}
}
