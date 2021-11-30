package MedianFinder

import "container/heap"

// 使用两个堆进行实现，一个大堆 一个小堆
// 如果总数是奇数  则取出较长堆中的顶点值
// 大根堆在前 小根堆在后

type bighp []int

func (b *bighp) Less(i, j int) bool {
	return (*b)[i] > (*b)[j]
}
func (b *bighp) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}
func (b *bighp) Len() int {
	return len(*b)
}
func (b *bighp) Push(x interface{}) {
	*b = append(*b, x.(int))
}
func (b *bighp) Pop() interface{} {
	x := (*b)[b.Len()-1]
	old := (*b)
	*b = (*b)[:len(old)-1]
	return x
}

type smallhp []int

func (s *smallhp) Less(i, j int) bool {
	return (*s)[i] < (*s)[j]
}
func (s *smallhp) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
func (s *smallhp) Len() int {
	return len(*s)
}
func (s *smallhp) Push(x interface{}) {
	*s = append(*s, x.(int))
}
func (s *smallhp) Pop() interface{} {
	x := (*s)[s.Len()-1]
	old := *s
	*s = (*s)[:len(old)-1]
	return x
}

type MedianFinder struct {
	left  bighp
	right smallhp
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := MedianFinder{
		left:  bighp{},
		right: smallhp{},
	}
	heap.Init(&(m.left))
	heap.Init(&(m.right))
	return m
}

func (f *MedianFinder) AddNum(num int) {
	// 由于是分两边
	// left存储较小的一段  right存储较大的一段
	// 如果 两个长度相同，则应该向存储较大的一段里面增加元素  则是先向大顶堆增加元素，然后再将大顶堆的堆顶元素加入到小顶堆
	// 如果长度不相同 则应该向存储较小的一段里面增加元素  则是先向小顶堆增加元素，然后再将小顶堆的堆顶元素加入到大顶堆中
	if len(f.left) != len(f.right) {
		heap.Push(&(f.right), num)
		top := heap.Pop(&(f.right))
		heap.Push(&(f.left), top)
	} else {
		heap.Push(&(f.left), num)
		top := heap.Pop(&(f.left))
		heap.Push(&(f.right), top)
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if len(f.left) != len(f.right) {
		//两者长度不相同 则为奇数
		return float64(f.right[0])
	} else {
		return (float64(f.left[0]) + float64(f.right[0])) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
