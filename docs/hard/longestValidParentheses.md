# [32. 最长有效括号](https://leetcode-cn.com/problems/longest-valid-parentheses/)

给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

> 输入: "(()"
>
> 输出: 2
>
> 解释: 最长有效括号子串为 "()"

### 解题思路

不能使用栈，因为可能出现()(()这种情况

使用动态规划，整个大前提是遇到左括号就设置为0，右括号分为两种情况

第一种情况: 正常情况 “......()” 即s[i] = ')' s[i-1] = '('，则:
$$
dp[i] = dp[i-2] + 2
$$
第二种情况，连续相同的括号:"...........))",即s[i] = ')' s[i-1] = ')'，则这种情况更加复杂，
$$
dp[i] = dp[i-1] + dp[i-dp[i-1]-1] +2 
$$
式子中，dp[i-1]就是上一个位置的最大长度，而dp[i-dp[i-1]-1]是组成括号之前是否还有已经组成的括号.

```go
func longestValidParentheses(s string) int {
	//使用栈  不行
	n := len(s)
	dp := make([]int, n)
	// 使用动态规划
	res := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			//两种情况
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		res = max(res, dp[i])

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

