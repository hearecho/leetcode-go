package SearchMatrix

import (
	"fmt"
	"testing"
)

type param struct {
	martix [][]int
	target int
}
func Test_searchMatrix(t *testing.T) {
	params := []param{
		{[][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},30},
	}
	for _,p:= range params{
		fmt.Printf("【input】:%v\t【output】:%v\n",p,searchMatrix(p.martix,p.target))
	}
}