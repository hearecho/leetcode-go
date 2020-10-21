package zuheshuzi2

import "sort"

//这个不可重复使用数字，但是数组里面是有重复数字的
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	//排序防止重复
	sort.Ints(candidates)
	backtrack(&res, make([]int, 0), 0, target, candidates)
	return res
}

func backtrack(res *[][]int, curArr []int, index int, target int, candidates []int) {
	if target <= 0 {
		if target == 0 {
			b := make([]int, len(curArr))
			copy(b, curArr)
			*res = append(*res, b)
		}
		return
	}
	//开始进行回溯
	for i := index; i < len(candidates); i++ {
		//去重 回溯去重中的关键语句，主要是适用于经过排序之后，相等元素在一块，并且排除第一次元素等于index的情况
		//这种情况只会出现一次
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		//剪枝
		if target < candidates[i] {
			return
		}
		curArr = append(curArr, candidates[i])
		//i+1的原因是因为不需要元素不能重复使用
		backtrack(res, curArr, i+1, target-candidates[i], candidates)
		//减去刚刚加在末尾的
		curArr = curArr[:len(curArr)-1]
	}
}
