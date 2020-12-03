package structure

import (
	"fmt"
	"testing"
)

func Test_Heap(t *testing.T)  {
	heap := NewMaxHeap([]int{7,8,5,9,1,2})
	heap.Peek()
	fmt.Println(heap)
}
