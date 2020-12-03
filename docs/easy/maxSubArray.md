# 最大子序和
> 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

### 注意点
1. 连续子数组，且子数组最少包含一个元素

### 解题思路
1. 判断当前总和的值是否小于0，如果小于等于0，则直接令综合等于当前遍历的值。

```go
func maxSubArray(nums []int) int {
	max := nums[0]
	cursum := 0
	for _, num := range nums {
		if cursum <= 0 {
			cursum = num
		} else {
			cursum += num
		}
		max = Max(max, cursum)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```