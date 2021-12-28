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
}

//单调队列
type dqueue []int

func (d *dqueue) Push(x int) {
	for len(*d) != 0 && (*d)[len(*d)-1] < x {
		// 移除前面比该值小的数
		*d = (*d)[:len(*d)-1]
	}
	*d = append(*d, x)
}

func (d *dqueue) Max() int {
	if len(*d) == 0 {
		return -1
	}
	return (*d)[0]
}

func (d *dqueue) Pop(n int) {
	if len(*d) != 0 && (*d)[0] == n {
		*d = (*d)[1:]
	}
}

func maxSlidingWindowQueue(nums []int, k int) []int {
	// 使用单调队列或者是 优先队列 优先队列存储的索引，但是按照的是最大值进行的排序
	res := make([]int, 0)
	i, j := 0, 0
	q := make(dqueue, 0)
	for j <= len(nums) {
		if j-i < k {
			q.Push(nums[j])
		} else {
			t := nums[i]
			res = append(res, q.Max())
			q.Pop(t)
			if j < len(nums) {
				q.Push(nums[j])
			}
			i++
		}
		j++
	}
	return res
}
