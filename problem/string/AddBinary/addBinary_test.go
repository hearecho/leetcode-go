package AddBinary

import (
	"fmt"
	"testing"
)

type param struct {
	a,b string
}

func Test_addBinary(t *testing.T)  {
	params := []param {
		{"11","1"},
		{"1010","1011"},
		{"0","0"},
	}
	for _,p := range params {
		fmt.Printf("【input】:%v\t【output】%v\n",p,addBinary(p.a,p.b))
	}
}