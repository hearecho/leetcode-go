package Median_of_two_sorted_arrays

import (
	"fmt"
	"testing"
)

type param struct {
	nums1 []int
	nums2 []int
}
func Test_div(t *testing.T)  {
	fmt.Println(5/2)
}

func Test_findMedianSortedArrays(t *testing.T)  {
	params := []param{
		{[]int{1,3},[]int{2}},
		{[]int{1,2},[]int{3,4}},
		{[]int{1},[]int{}},
		{[]int{1,4,8,9},[]int{2,4,4,4}},
	}
	for _,p := range params{
		fmt.Printf("【input】:%v       【output】:%v\n",p,findMedianSortedArrays(p.nums1,p.nums2))
	}
}
