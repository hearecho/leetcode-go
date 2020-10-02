package Add_Two_Numbers

import (
	"leetcode-go/problem/linkedlist/listNode"
)

func addTwoNumbers(l1 *listNode.ListNode, l2 *listNode.ListNode) *listNode.ListNode {
	head := &listNode.ListNode{}
	last := head
	carry := 0
	for l1 != nil && l2 != nil {
		num := l1.Val + l2.Val + carry
		carry = num / 10
		temp := listNode.ListNode{
			Val:  num % 10,
			Next: nil,
		}
		//尾插法
		last.Next = &temp
		last = &temp
		l1 = l1.Next
		l2 = l2.Next
	}
	//处理l1 ,l2剩余的结果
	for l1 != nil {
		num := l1.Val + carry
		carry = num / 10
		temp := listNode.ListNode{
			Val:  num % 10,
			Next: nil,
		}
		//尾插法
		last.Next = &temp
		last = &temp
		l1 = l1.Next
	}
	for l2 != nil {
		num := l2.Val + carry
		carry = num / 10
		temp := listNode.ListNode{
			Val:  num % 10,
			Next: nil,
		}
		//尾插法
		last.Next = &temp
		last = &temp
		l2 = l2.Next
	}
	if carry != 0 {
		temp := listNode.ListNode{
			Val:  1,
			Next: nil,
		}
		last.Next = &temp
		last = &temp
	}

	return head.Next
}

func optimize_addTwoNumbers(l1 *listNode.ListNode, l2 *listNode.ListNode) *listNode.ListNode {
	head := &listNode.ListNode{}
	last := head
	carry := 0
	n1,n2 := 0,0
	for l1!=nil || l2!= nil || carry != 0{
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		temp := &listNode.ListNode{
			Val:  (n1+n2+carry)%10,
			Next: nil,
		}
		last.Next = temp
		last = temp
		carry =( n1+n2+carry)/10
	}
	return head.Next
}
