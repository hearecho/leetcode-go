package reverseKList

import (
	"leetcode-go/problem/linkedlist/listNode"
	"testing"
)

func Test_reverseKGroup(t *testing.T)  {
	head := listNode.GenerateLinkedListTail([]int{1,2,3,4,5})
	res := reverseKGroup(head,3)
	listNode.PrinLinkedList(res)
}
