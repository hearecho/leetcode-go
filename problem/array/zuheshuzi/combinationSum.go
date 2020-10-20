package zuheshuzi

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	combinations := make([][]int, 0)
	sort.Ints(candidates)
	backtrack(make([]int, 0), target, &combinations, candidates, 0)
	return combinations
}

func backtrack(curArr []int, target int, combinations *[][]int, candidates []int, index int) {
	if target == 0 {
		b := make([]int, len(curArr))
		copy(b, curArr)
		*combinations = append(*combinations, b)
		return
	}
	if target < 0 {
		return
	}

	//开始向下减
	for i := index; i < len(candidates); i++ {
		//剪纸
		if target < candidates[i] {
			return
		}
		curArr = append(curArr, candidates[i])
		fmt.Printf("【cuarr】%v\ttarget %v\n", curArr, target-candidates[i])
		backtrack(curArr, target-candidates[i], combinations, candidates, i)
		curArr = curArr[:len(curArr)-1]
	}
}
