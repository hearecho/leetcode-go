package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

// 利用递归的思想
func reverseList(head *ListNode) *ListNode {
	// 使用头插法就可以完成反转
	// 也可以使用递归
	if head == nil || head.Next == nil {
		return head
	}
	// 以链表 1->2->3->4->5 为例，现在链表变成了 5->4->3->2->null，1->2->null（是一个链表，不是两个链表）
	// 此时 newHead是5，head是1
	new := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return new
}
