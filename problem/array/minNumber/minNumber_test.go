package minNumber

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	nums := sortNums{3, 30, 34, 5, 9}
	sort.Sort(nums)
	fmt.Println(nums)
}

func TestFunc(t *testing.T) {
	nums := []int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}
	fmt.Println(minNumber(nums))
}
