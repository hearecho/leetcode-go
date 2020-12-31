package dailyProblem

import (
	"sort"
)

//https://leetcode-cn.com/problems/non-overlapping-intervals/

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	res, end := 0, intervals[0][1]
	//一旦重叠就删除该区间
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] < end {
			continue
		}
		end = intervals[i][1]
		res++
	}
	return len(intervals) - res
}
