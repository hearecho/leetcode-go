package contest210

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T)  {
	fmt.Println(slowestKey([]int{9,29,49,50},"cbcd")+'0')
}

func Test_2(t *testing.T)  {
	fmt.Println(checkArithmeticSubarrays([]int{4,6,5,9,3,7},[]int{0,0,2},[]int{2,3,5}))
}

func Test_3(t *testing.T)  {
	fmt.Println(minimumEffortPath([][]int{{1,2,3},{3,8,4},{5,3,5}}))
}

func Test_4(t *testing.T)  {

}
