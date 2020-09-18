package structure

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateLinkedListTail(nums []int) *ListNode {
	head := &ListNode{}
	last := head
	//尾插法构建
	for _,num := range nums {
		temp := &ListNode{
			Val:  num,
			Next: nil,
		}
		last.Next = temp
		last = temp
	}
	return head.Next
}

func PrinLinkedList(l *ListNode)  {
	for l != nil {
		fmt.Printf("%d\t",l.Val)
		l = l.Next
	}
}
