package dailyProblem

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	//按照宽的大小进行排序，如果宽相同，则按照高度从小到大的顺序进行排序，寻找刚好装下的那一个
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	//和最长上升子序列，只不过这个是二维的 所以还是先排序 而且不在同一个宽度的区域进行两次选择
	//所以考虑使用 动态规划dp[i][j]表示从i~j 位置最多有几个信封
	dp := make([]int, len(envelopes))
	dp[0] = 1
	res := 1
	for i := 1; i < len(envelopes); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] && dp[j] > max {
				max = dp[j]
			}
		}
		if max == 0 {
			dp[i] = 1
		} else {
			dp[i] = max + 1
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
