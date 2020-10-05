package mergeKLinkedlist

import (
	"fmt"
	"leetcode-go/problem/linkedlist/listNode"
	"testing"
)

type param struct {
	nums [][]int
}

func Test_mergeKLists(t *testing.T)  {
	params := []param{
		{[][]int{{1,4,5},{1,3,4},{2,6}}},
	}
	for _,p := range params {
		lists := []*listNode.ListNode{}
		for _,num := range p.nums {
			lists = append(lists,listNode.GenerateLinkedListTail(num))
		}
		res := mergeKLists(lists)
		fmt.Printf("【input】:%v      【output】:",p)
		listNode.PrinLinkedList(res)
		fmt.Println()
	}
}