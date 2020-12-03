package MergeTwoLists

import "leetcode-go/problem/linkedlist/listNode"

func mergeTwoLists(l1 *listNode.ListNode, l2 *listNode.ListNode) *listNode.ListNode {
	fakeHead := &listNode.ListNode{}
	last := fakeHead
	for l1 != nil || l2 != nil {
		if l1 == nil {
			last.Next = l2
			break
		}
		if l2 == nil {
			last.Next = l1
			break
		}
		if l1.Val <= l2.Val {
			last.Next = l1
			last = l1
			l1 = l1.Next
		} else {
			last.Next = l2
			last = l2
			l2 = l2.Next
		}
	}
	return fakeHead.Next
}