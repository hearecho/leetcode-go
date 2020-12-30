package dailyProblem

import (
	"sort"
)

//https://leetcode-cn.com/problems/last-stone-weight/

func lastStoneWeight(stones []int) int {
	sort.Slice(stones, func(i, j int) bool {
		return stones[i] > stones[j]
	})
	for len(stones) > 1 {
		num1, num2 := stones[0], stones[1]
		stones = stones[2:]
		if num1 != num2 {
			stones = append(stones, num1-num2)
		}
		sort.Slice(stones, func(i, j int) bool {
			return stones[i] > stones[j]
		})
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
