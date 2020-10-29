# 跳跃游戏
> 给定一个非负整数数组，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。
> 判断你是否能够到达最后一个位置。

### 注意点
> 元素代表你在该位置可以跳跃的最大长度

### 解题思路
> 我们的目标是最后可以到达最后的目的地，所以，我们可以查找每一次跳跃能够到达最远的地方，即`next_indx+nums[next_indx] <= nums[i+j]+i+j`其中next_index是下一个位置的index，i+j则是当前位置可能会跳跃到下一个位置的index，如果成功则应该更新next_index的值。主要使用的就是贪心算法，每次都找能够跳跃的最大位置。

```go
func canJump(nums []int) bool {
	l := len(nums)
	for i, next_indx := 0, 0; i < l; {
		if next_indx != l-1 && nums[next_indx] == 0 {
			return false
		}
		if next_indx == l-1 {
			return true
		}
		next_indx++
		for j := 1; j <= nums[i]; j++ {
			if i+j < l && next_indx+nums[next_indx] <= nums[i+j]+i+j {
				next_indx = i + j
			}
		}
		i = next_indx
	}
	return false
}
```