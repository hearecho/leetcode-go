package Mypow

import (
	"fmt"
	"testing"
)

type param struct {
	x float64
	n int
}

func Test_Mypow(t *testing.T)  {
	params := []param{
		{2.00000,10},
		{2.10000,3},
		{2.00000,-2},
	}
	for _,p := range params {
		fmt.Printf("【input】:%v\t【ouput】:%v\n",p,myPow(p.x,p.n))
	}
}