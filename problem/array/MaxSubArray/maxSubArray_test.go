package MaxSubArray

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
}

func Test_maxSubArray(t *testing.T) {
	params := []param{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, maxSubArray(p.nums))
	}
}
