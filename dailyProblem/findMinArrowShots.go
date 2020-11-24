package dailyProblem

import "sort"

/**
https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/
*/
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	res := 1
	//实际题目就是查询重叠区间的个数，但是每次重叠都是要收紧区间的大小，以便能够使用一支箭就能射穿全部的气球
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	h := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= h {
			//有重合部分,那就收紧
			h = min(h, points[i][1])
			continue
		}
		h = points[i][1]
		res++
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
