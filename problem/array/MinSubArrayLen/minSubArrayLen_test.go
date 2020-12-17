package MinSubArrayLen

import (
	"fmt"
	"testing"
)

type param struct {
	nums []int
	s    int
}

func Test_minSubArrayLen(t *testing.T) {
	params := []param{
		{[]int{2, 3, 1, 2, 4, 3}, 7},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, minSubArrayLen(p.s, p.nums))
	}
}
