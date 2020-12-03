package NthFromEnd

import (
	"leetcode-go/problem/linkedlist/listNode"
)

//双指针加上假的头节点
func removeNthFromEnd(head *listNode.ListNode, n int) *listNode.ListNode {
	fakehead := &listNode.ListNode{Val: 0}
	fakehead.Next = head
	//设置两个指针，第一个指针指向头部，第二个指针指向第n个位置
	low,fast:= fakehead,fakehead
	for n >= 0 {
		fast = fast.Next
		n--
	}
	//一起循环直到fast到达末尾，然后删除low指向的节点
	for fast != nil {
		fast = fast.Next
		low = low.Next
	}
	//删除low的下一个节点
	temp := low.Next
	low.Next = temp.Next
	return fakehead.Next
}
