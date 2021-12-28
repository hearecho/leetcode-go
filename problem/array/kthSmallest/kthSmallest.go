package kthSmallest

import "container/heap"

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

// 构造一个大根堆 然后进行遍历即可
func kthSmallest(matrix [][]int, k int) int {
	hp := make(bighp, 0)
	heap.Init(&hp)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if len(hp) < k {
				heap.Push(&hp, matrix[i][j])
			} else if hp[0] > matrix[i][j] {
				// 需要取出堆顶元素
				heap.Pop(&hp)
				heap.Push(&hp, matrix[i][j])
			}
		}
	}
	return hp[0]
}
