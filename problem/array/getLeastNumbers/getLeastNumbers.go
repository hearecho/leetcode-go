package getLeastNumbers

import "container/heap"

func getLeastNumbers(arr []int, k int) []int {
	// 快排 判断快排的的index
	// 使用堆
	var quickSort func(l, r int)
	var partion func(l, r int) int
	quickSort = func(l, r int) {
		if l >= r {
			return
		}
		i := partion(l, r)
		quickSort(l, i-1)
		quickSort(i+1, r)
	}
	partion = func(l, r int) int {
		i, j := l, r
		for i < j {
			// 必须从后往前遍历
			for i < j && arr[j] >= arr[l] {
				j--
			}
			for i < j && arr[i] <= arr[l] {
				i++
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[l], arr[i] = arr[i], arr[l]
		return i
	}
	quickSort(0, len(arr)-1)
	return arr[:k]
}

// 第二种方法是利用 快速排序 快速排序的过程中比较 每次切分数组的位置
// 返回结果没有要求有序，所以只需要k个最小的数字返回即可
// 这种方法就相当于是利用剪枝的方式 减小递归深度 提前结束递归
func getLeastNumbersNotSort(arr []int, k int) []int {
	// 快排 判断快排的的index
	// 使用堆
	var quickSort func(l, r int)
	var partion func(l, r int) int
	quickSort = func(l, r int) {
		if l >= r {
			return
		}
		i := partion(l, r)
		if i == k {
			return
		}
		if i > k {
			// i 如果比k大那么 这个时候需要继续划分左区间
			quickSort(l, i-1)
		} else {
			quickSort(i+1, r)
		}
	}
	partion = func(l, r int) int {
		i, j := l, r
		for i < j {
			// 必须从后往前遍历
			for i < j && arr[j] >= arr[l] {
				j--
			}
			for i < j && arr[i] <= arr[l] {
				i++
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[l], arr[i] = arr[i], arr[l]
		return i
	}
	quickSort(0, len(arr)-1)
	return arr[:k]
}

// 使用堆（优先队列）的方式
// 创建一个大小为k的大根堆
// 如果堆的空间已经满了，则与堆顶值进行比较 如果小于则抛弃堆顶并入堆 大于则不做判断

type hp []int

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func getLeastNumbersByHeap(arr []int, k int) []int {
	h := &hp{}
	heap.Init(h)
	if k == 0 {
		return []int{}
	}
	for i := 0; i < len(arr); i++ {
		if len(*h) < k {
			heap.Push(h, arr[i])
		} else {
			if (*h)[0] > arr[i] {
				heap.Pop(h)
				heap.Push(h, arr[i])
			}
		}
	}
	return *h
}
