package MergeInterval

import (
	"fmt"
	"testing"
)

type param struct {
	intervals [][]int
}

func Test_mergeIntervals(t *testing.T) {
	params := []param{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
	}
	for _, p := range params {
		temp := p.intervals
		fmt.Printf("【input】:%v\t【output】:%v\n", temp, merge(p.intervals))
	}
}
