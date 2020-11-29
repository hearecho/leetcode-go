package dailyProblem

import "sort"

//https://leetcode-cn.com/problems/largest-perimeter-triangle/
func largestPerimeter(A []int) int {
	//面积不为0，则是说明A可以组成三角形
	sort.Ints(A)
	for i := len(A) - 1; i >= 2; i-- {
		if A[i-2]+A[i-1] > A[i] {
			return A[i-2] + A[i-1] + A[i]
		}
	}
	return 0
}
