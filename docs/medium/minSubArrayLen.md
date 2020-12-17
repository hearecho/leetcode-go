# [209. 长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)



给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

> ```
>输入：s = 7, nums = [2,3,1,2,4,3]
>输出：2
>解释：子数组 [4,3] 是该条件下的长度最小的子数组。
> ```

### 解题思路

> 双指针思想（滑动窗口），主要是维护一个滑动窗口，之后如果大于则减少，小于则继续增加窗口大小。

```go
func minSubArrayLen(s int, nums []int) int {
    // 滑动窗口思想
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, 0
	curSum := 0
	ans := math.MaxInt32
	for r < len(nums) {
		curSum += nums[r]
		for curSum >= s {
			ans = min(r-l+1, ans)
			curSum -= nums[l]
			l++
		}
		r++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

