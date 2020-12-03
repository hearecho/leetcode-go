package reverseKList

import "leetcode-go/problem/linkedlist/listNode"

/**
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
想法:每k个链表截取一端,之后反转整个链表。如果无法截取k个节点则不进行反转,之后进行尾插,
可以返回两个节点，一个是头，一个是尾，便于定位
 */
func reverseKGroup(head *listNode.ListNode, k int) *listNode.ListNode {
	fakehead := &listNode.ListNode{}
	last := fakehead
	for head != nil {
		i := 0
		temp:= head
		for i<k-1 && head != nil{
			head = head.Next
			i++
		}
		if i < k-1 || head == nil{
			last.Next = temp
			break
		}
		next := head.Next
		//断开
		head.Next = nil
		last.Next = reverseList(temp)
		last = temp
		head = next
	}
	return fakehead.Next
}

/**
反转链表只需要进行头插法就好
 */
func reverseList(head *listNode.ListNode) *listNode.ListNode  {
	fakehead := &listNode.ListNode{}
	for head != nil {
		next := head.Next
		//头插
		head.Next = fakehead.Next
		fakehead.Next = head
		head = next
	}
	return fakehead.Next
}
