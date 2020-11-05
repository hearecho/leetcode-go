package UniquePathsWithObstacles

import (
	"fmt"
	"testing"
)

type param struct {
	grids [][]int
}

func Test_uniquePathsWithObstacles(t *testing.T)  {
	params := []param{
		{[][]int{{0,0,0},{0,1,0},{0,0,0}}},
		{[][]int{{0}}},
		{[][]int{{1,0}}},
	}
	for _,p := range  params {
		fmt.Printf("【input】:%v\t【output】:%v\n",p,uniquePathsWithObstacles(p.grids))
	}
}