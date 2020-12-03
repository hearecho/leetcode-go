package findMissNum

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
}

func Test_firstMissingPositive(t *testing.T) {
	params := []param{
		{[]int{1, 2, 0}},
		{[]int{3, 4, -1, 1}},
		{[]int{7, 8, 9, 11, 12}},
	}
	for _, p := range params {
		fmt.Printf("【input】%v\t【output】%v\n", p, firstMissingPositive(p.nums))
	}
}
