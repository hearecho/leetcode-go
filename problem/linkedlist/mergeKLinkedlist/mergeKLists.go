package mergeKLinkedlist

import "leetcode-go/problem/linkedlist/listNode"

//合并两个是合并
//依次合并即可
func mergeKLists(lists []*listNode.ListNode) *listNode.ListNode {
	temphead := &listNode.ListNode{}
	for i:=0;i<len(lists);i++ {
		temphead.Next = mergeTwoLists(temphead.Next,lists[i])
	}
	return temphead.Next
}

//两两合并
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

//使用归并排序的方式进行合并,这个类似归并排序已经分成了几段，我们可以继续分
//将每个链表看作是一个数，这样后续合并也是合并的参数
func mergeKLists_optimaize1(lists []*listNode.ListNode) *listNode.ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists(left, right)
}

//可以将所有的元素的都遍历然后存储在一个小根堆中，之后再取出