package VideoStitching

import (
	"math"
	"sort"
)

func videoStitching(clips [][]int, T int) int {
	//按照开始节点paixue
	sort.Slice(clips[:], func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})
	if clips[0][0] != 0 {
		return -1
	}
	max := clips[0][1]
	count := 1
	for i := 1; i < len(clips); i++ {
		if clips[i][0] == 0 && clips[i][1] > max {
			max = clips[i][1]
		}
		if clips[i][0] < max && clips[i][1] > max {
			max = clips[i][1]
			count++
		}
	}
	if max < T {
		return -1
	}
	return count
}

func videoStitching_1(clips [][]int, t int) int {
	const inf = math.MaxInt64 - 1
	dp := make([]int, t+1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 1; i <= t; i++ {
		for _, c := range clips {
			l, r := c[0], c[1]
			// 若能剪出子区间 [l,i]，则可以从 dp[l] 转移到 dp[i]
			if l < i && i <= r && dp[l]+1 < dp[i] {
				dp[i] = dp[l] + 1
			}
		}
	}
	if dp[t] == inf {
		return -1
	}
	return dp[t]
}

func videoStitching_2(clips [][]int, t int) (ans int) {
	maxn := make([]int, t)
	last, pre := 0, 0
	for _, c := range clips {
		l, r := c[0], c[1]
		if l < t && r > maxn[l] {
			maxn[l] = r
		}
	}
	for i, v := range maxn {
		if v > last {
			last = v
		}
		if i == last {
			return -1
		}
		if i == pre {
			ans++
			pre = last
		}
	}
	return
}
