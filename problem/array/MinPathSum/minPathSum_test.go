package MinPathSum

import (
	"fmt"
	"testing"
)

type param struct {
	grid [][]int
}

func Test_MinPathSum(t *testing.T)  {
	params := []param{
		{[][]int{{1,3,1},{1,5,1},{4,2,1}}},
	}

	for _,p := range  params {
		fmt.Printf("【input】:%v\t【output】:%v\n",p,minPathSum(p.grid))
	}
}