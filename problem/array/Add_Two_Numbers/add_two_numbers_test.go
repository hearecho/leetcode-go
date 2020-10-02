package Add_Two_Numbers

import (
	"fmt"
	"leetcode-go/problem/linkedlist/listNode"
	"testing"
)

type params struct {
	l1 *listNode.ListNode
	l2 *listNode.ListNode
}
func Test_addTwoNumbers(t *testing.T)  {
	ps := []params{
		{listNode.GenerateLinkedListTail([]int{305}), listNode.GenerateLinkedListTail([]int{4,5,6})},
		{listNode.GenerateLinkedListTail([]int{1}), listNode.GenerateLinkedListTail([]int{9,9,9})},
		{listNode.GenerateLinkedListTail([]int{0,4}), listNode.GenerateLinkedListTail([]int{8,0,1})},
	}
	for _,p := range ps{
		fmt.Printf("【input】:%v       【output】:%v\n",p,*addTwoNumbers(p.l1,p.l2))
	}
}
