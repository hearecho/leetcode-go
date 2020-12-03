package romanToInt

import (
	"fmt"
	"testing"
)

type param struct {
	s string
}

func Test_RomanToInt(t *testing.T)  {
	params := []param{
		{"III"},
		{"IV"},
		{"IX"},
		{"LVIII"},
		{"MCMXCIV"},
	}

	for _,p := range params {
		fmt.Printf("【input】:%v       【output】:%v\n",p,optimaize_romanToInt(p.s))
	}
}
