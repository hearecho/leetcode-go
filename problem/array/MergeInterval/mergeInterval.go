package MergeInterval

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return res
	}
	//排序按照start端点，从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	curArr := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > curArr[1] {
			//证明不可再继续合并
			b := make([]int, 2)
			copy(b, curArr)
			res = append(res, b)
			curArr = intervals[i]
		} else {
			//合并
			if intervals[i][1] > curArr[1] {
				curArr[1] = intervals[i][1]
			}
		}
	}
	res = append(res, curArr)
	return res

}

//插入区间的题解
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)

}
