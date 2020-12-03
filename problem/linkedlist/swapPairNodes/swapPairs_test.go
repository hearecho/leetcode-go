package swapPairNodes

import (
	"fmt"
	"leetcode-go/problem/linkedlist/listNode"
	"testing"
)

type param struct {
	nums []int
}

func Test_swapPairs(t *testing.T)  {
	params := []param{
		{[]int{2,1,4,3,6,5}},
	}

	for _,p := range params {
		head := listNode.GenerateLinkedListTail(p.nums)
		res := swapPairs(head)
		fmt.Printf("【input】:%v\n       【output】:\n",p)
		listNode.PrinLinkedList(res)
		fmt.Println()
	}
}
