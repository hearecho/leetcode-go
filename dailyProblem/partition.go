package dailyProblem

//https://leetcode-cn.com/problems/partition-list/

func partition(head *ListNode, x int) *ListNode {
	//直接新构建两个链表之后将两个链表进行合并,这样顺序还不会发生变化
	//使用尾插法
	less, more := &ListNode{}, &ListNode{}
	l1, l2 := less, more
	for head != nil {
		temp := head.Next
		head.Next = nil
		if head.Val < x {
			l1.Next = head
			l1 = head
		} else {
			l2.Next = head
			l2 = head
		}
		head = temp
	}
	l1.Next = more.Next
	return less.Next
}
