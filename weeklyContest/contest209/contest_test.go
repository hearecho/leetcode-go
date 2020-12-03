package contest209

import (
	"fmt"
	"testing"
)

func Test_contest1(t *testing.T)  {
	fmt.Println(specialArray([]int{0,4,3,0,4}))
}

func Test_contest3(t *testing.T)  {
	res := visiblePoints([][]int{{0,0},{0,2}},90,[]int{1,1})

	fmt.Println(res)
}
