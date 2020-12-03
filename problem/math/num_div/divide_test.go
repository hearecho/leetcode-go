package num_div

import (
	"fmt"
	"testing"
)

type param struct {
	dividend int
	divisor int
}

func Test_divide(t *testing.T)  {
	params := []param{
		{10,3},
		{7,-3},
	}
	for _,p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n",p,true_divide(p.dividend,p.divisor))
	}
}
