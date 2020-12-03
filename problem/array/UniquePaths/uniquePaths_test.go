package UniquePaths

import (
	"fmt"
	"testing"
)

type param struct {
	m,n int
}

func Test_UniquePaths(t *testing.T)  {
	params := []param {
		{3,2},
		{7,3},
	}
	for _,p := range params {
		fmt.Printf("【iutput】:%v\t【output】%v\n",p,uniquePaths(p.m,p.n))
	}
}