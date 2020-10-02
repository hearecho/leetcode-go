package NthFromEnd

import (
	"fmt"
	"leetcode-go/problem/linkedlist/listNode"
	"testing"
)

type param struct {
	nums []int
	n    int
}

func Test_removeNthFromEnd(t *testing.T)  {
	params := []param{
		{[]int{1,2,3,4,5},2},
	}
	for _,p := range params {
		head := listNode.GenerateLinkedListTail(p.nums)
		res := removeNthFromEnd(head,p.n)
		fmt.Printf("【input】:%v\n       【output】:\n",p)
		listNode.PrinLinkedList(res)
		fmt.Println()
	}
}
