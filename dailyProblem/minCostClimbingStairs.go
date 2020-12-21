package dailyProblem

//https://leetcode-cn.com/problems/min-cost-climbing-stairs/
func minCostClimbingStairs(cost []int) int {
	//动态规划
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}
