package MergeTwoLists

import (
	"fmt"
	"leetcode-go/problem/linkedlist/listNode"
	"testing"
)

type param struct {
	nums1 []int
	nums2 []int
}
func Test_Merge(t *testing.T)  {
	params := []param {
		{[]int{1,2,4},[]int{1,3,4}},
		{[]int{1,2,3},[]int{1,8,9}},
		{[]int{1,},[]int{0,5,6}},
	}
	for _,p := range params {
		l1 := listNode.GenerateLinkedListTail(p.nums1)
		l2 := listNode.GenerateLinkedListTail(p.nums2)
		res := mergeTwoLists(l1,l2)
		fmt.Printf("【input】:%v      【output】:",p)
		listNode.PrinLinkedList(res)
		fmt.Println()
	}
}