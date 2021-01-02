package dailyProblem

import (
	"container/heap"
	"sort"
)

//https://leetcode-cn.com/problems/sliding-window-maximum/

//使用优先队列 最大堆就可以进
//go 创建最大堆
type hp struct {
	sort.IntSlice
}

var a []int

//实现Less方法来表示最大还是最小堆
func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}
func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	h := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		h.IntSlice[i] = i
	}
	heap.Init(h)
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[h.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(h, i)
		for h.IntSlice[0] <= i-k {
			heap.Pop(h)
		}
		ans = append(ans, nums[h.IntSlice[0]])
	}
	return ans
}
