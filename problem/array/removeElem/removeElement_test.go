package removeElem

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
	val  int
}

func Test_removeElement(t *testing.T) {
	params := []param{
		{[]int{3, 2, 2, 3},3},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2},2},
	}
	for _, p := range params {
		fmt.Printf("【nums】:%v       【output】:%v\n", p, removeElement(p.nums,p.val))
	}
}
