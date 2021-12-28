package wordBreak

func wordBreak(s string, wordDict []string) bool {
	// 使用动态规划的方式
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	// dp[i] 表示 s[0:i-1]是否在
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			dp[i] = dp[j] && wordDictSet[s[j:i]]
			if dp[i] {
				break
			}
		}
	}
	return dp[len(s)]
}

func wordBreak2(s string, wordDict []string) []string {
	// 相比于上一个 这一个需要对所有可组成的单词进行记录
	// 因为是多段  使用回溯
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	res := make([]string, 0)

	var dfs func(start int, curStr string)

	dfs = func(start int, curStr string) {
		if start == len(s) {
			res = append(res, curStr[1:])
		}
		for i := start + 1; i <= len(s); i++ {
			if wordDictSet[s[start:i]] {
				dfs(i, curStr+" "+s[start:i])
			}
		}
	}
	dfs(0, "")
	return res
}
