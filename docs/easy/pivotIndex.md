# [寻找数组中心索引](https://leetcode-cn.com/leetbook/read/array-and-string/yf47s/)
> 给定一个整数类型的数组 nums，请编写一个能够返回数组 “中心索引” 的方法。
  我们是这样定义数组 中心索引 的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
  如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。

### 解题思路
> 只需要知道全部和和左边和就可以知道右边和
```go
func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	left := 0
	for i := 0; i < len(nums); i++ {
		// right = sum-num[i]-left
		if left == sum-nums[i]-left {
			return i
		}
		left += nums[i]
	}
	return -1
}

```