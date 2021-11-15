package maxSlidingWindow

import (
	"container/heap"
	"math"
	"sort"
)

// 最大堆和最小堆的区别就是 less方法的区别
var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	// 堆的构建使用到了数组，所以堆主要存储的是
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		// 顶部存储的是当前最大值的index值 i-k是被删除序列的值
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

func maxSlidingWindowBaoli(nums []int, k int) []int {
	// 暴力法
	res := []int{}
	for i := 0; i <= len(nums)-k; i++ {
		max := math.MinInt64
		for j := i; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		res = append(res, max)
	}
	return res

	// 使用最大堆的方式
	// 维持一个大小为k的最大堆

}
