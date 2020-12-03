package structure

import "fmt"

type Heap interface {
	Push(int)
	Peek()(int,error)
}
//堆的实现
//堆并不是真正意义的树，它实际上还是一个数组
//堆的结构
type minHeap struct {
	Elems []int
	Length int
}

//一个节点为i（从0开始计算）
//父节点 parent(i) = floor(i-1)/2
//左右子节点 left(i) = 2i+1 右节点 right(i) = 2i+2
//大顶堆的性质，在大顶堆中,父节点的值要大于或者等于左右子节点
/**
ty: 类型 0为大顶堆，1为小顶堆
 */
func NewMinHeap(nums []int) Heap {
	heap := &minHeap{
		Elems:  []int{},
		Length: 0,
	}
	for _,num := range nums {
		heap.Push(num)
	}
	return heap
}
//插入是从末尾插入,之后向上调整整个堆
func (h *minHeap)Push(elem int)  {
	//插入末尾
	h.Elems = append(h.Elems,elem)
	h.Length++
	//向上调整
	i := h.Length -1
	//如果父节点大于当前的值，则上浮(小顶堆)
	//i == 0证明已经到根了
	for h.Elems[(i-1)/2]>elem && i!=0 {
		h.Elems[i] = h.Elems[(i-1)/2]
		i = (i-1)/2
	}
	//更新此时i位置的值
	h.Elems[i] = elem
}

func (h *minHeap)Peek() (int,error)  {
	if h.Length == 0 {
		return 0,fmt.Errorf("heap is empty")
	}
	res := h.Elems[0]
	//向下调整整个最小堆
	lastElem := h.Elems[h.Length-1]
	i,child := 0,0
	for i*2+1 < h.Length {
		//左节点
		child = 2*i+1
		//右节点的数值更小
		if child < h.Length-1 && h.Elems[child+1] < h.Elems[child] {
			child++
		}
		if lastElem > h.Elems[child] {
			h.Elems[i] = h.Elems[child]
		} else {
			break
		}
		i = child
	}
	h.Elems[i] = lastElem
	h.Elems = h.Elems[:h.Length-1]
	h.Length--
	return res,nil
}

type maxHeap struct {
	Elems []int
	Length int
}
func NewMaxHeap(nums []int) Heap {
	heap := &maxHeap{
		Elems:  []int{},
		Length: 0,
	}
	for _,num := range nums {
		heap.Push(num)
	}
	return heap
}

func (h *maxHeap)Push(elem int)  {
	//插入末尾
	h.Elems = append(h.Elems,elem)
	h.Length++
	//向上调整
	i := h.Length -1
	//如果父节点大于当前的值，则上浮(小顶堆)
	//i == 0证明已经到根了
	for h.Elems[(i-1)/2]<elem && i!=0 {
		h.Elems[i] = h.Elems[(i-1)/2]
		i = (i-1)/2
	}
	//更新此时i位置的值
	h.Elems[i] = elem
}

func (h *maxHeap)Peek()(int,error)  {
	if h.Length == 0 {
		return 0,fmt.Errorf("heap is empty")
	}
	res := h.Elems[0]
	//向下调整整个最小堆
	lastElem := h.Elems[h.Length-1]
	i,child := 0,0
	for i*2+1 < h.Length {
		//左节点
		child = 2*i+1
		//右节点的数值更小
		if child < h.Length-1 && h.Elems[child+1] > h.Elems[child] {
			child++
		}
		if lastElem < h.Elems[child] {
			h.Elems[i] = h.Elems[child]
		} else {
			break
		}
		i = child
	}
	h.Elems[i] = lastElem
	h.Elems = h.Elems[:h.Length-1]
	h.Length--
	return res,nil
}


