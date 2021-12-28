package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{4, 5, 2, 6, 8, 4, 1, 2, 3}
	quickSort(nums)
	fmt.Println(nums)
}
