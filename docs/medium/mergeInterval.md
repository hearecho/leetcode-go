# [合并区间](https://leetcode-cn.com/problems/merge-intervals)
> 给出一个区间的集合，请合并所有重叠的区间。区间 [1,4] 和 [4,5] 可被视为重叠区间。

### 注意点
> 相等端点视为可合并端点

### 解题思路
> 使用贪心算法，将所有的区间按照起始节点从小到大排序，之后遍历整个数组，如果遍历的区间的起始节点大于当前结果区间的末尾节点，证明无法重合，添加到结果中，并将当并设置当前遍历区间为当前结果区间，如果小于则判断区间是否增大，如果增大更新。

```go
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return res
	}
	//排序按照start端点，从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	curArr := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > curArr[1] {
			//证明不可再继续合并
			b := make([]int, 2)
			copy(b, curArr)
			res = append(res, b)
			curArr = intervals[i]
		} else {
			//合并
			if intervals[i][1] > curArr[1] {
				curArr[1] = intervals[i][1]
			}
		}
	}
	res = append(res, curArr)
	return res
}
```
