package swapPairNodes

import "leetcode-go/problem/linkedlist/listNode"

//交换两个节点,那么节点就需要一次往后面进两步
func swapPairs(head *listNode.ListNode) *listNode.ListNode {
	fakehead := &listNode.ListNode{}
	last := fakehead
	for head != nil && head.Next != nil {
		next := head.Next.Next
		//交换两个节点
		last.Next = head.Next
		last.Next.Next = head
		//尾插法
		last = last.Next.Next
		head = next
	}
	last.Next = head
	return fakehead.Next
}
