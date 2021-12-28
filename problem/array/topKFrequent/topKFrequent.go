package topKFrequent

import "container/heap"

var counts map[int]int

type smallhp []int

func (s *smallhp) Less(i, j int) bool {
	return counts[(*s)[i]] < counts[(*s)[j]]
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

func topKFrequent(nums []int, k int) []int {
	// 哈希表存储 然后根据数量排序
	counts = make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	// 再使用堆 由于是最高频的元素 所以使用根据value值建立小根堆
	hp := make(smallhp, 0)
	heap.Init(&hp)
	for key, _ := range counts {
		if len(hp) < k {
			heap.Push(&hp, key)
		} else if counts[key] > counts[hp[0]] {
			heap.Pop(&hp)
			heap.Push(&hp, key)
		}
	}
	return hp
}
