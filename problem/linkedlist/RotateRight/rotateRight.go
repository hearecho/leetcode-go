package RotateRight

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	newHead := &ListNode{0,head}
	l := 0
	cur := newHead
	for cur.Next != nil {
		l++
		cur = cur.Next
	}
	k = k%l
	//如果k刚好是l的整数倍，那样的话，肯定会旋转为原来的样子,所以直接返回原字符串
	if k  == 0 {
		return head
	}
	cur.Next = head
	cur = newHead
	//找到中间
	for i := l - k; i > 0; i-- {
		cur = cur.Next
	}
	res := &ListNode{Val: 0, Next: cur.Next}
	cur.Next = nil
	return res.Next
}

type ListNode struct {
	Val int
	Next *ListNode
}