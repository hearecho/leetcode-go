package dailyProblem


func insertionSortList(head *ListNode) *ListNode {
	fakehead := &ListNode{}
	//双指针
	for head != nil {
		//获取当前的节点
		next := head.Next
		temp := head
		temp.Next = nil
		head = next
		//遍历已排序然后插入,找到大于temp节点值的节点的前一个节点
		cur := fakehead
		for cur.Next != nil {
			if cur.Next.Val > temp.Val {
				//将temp插入
				n := cur.Next
				cur.Next = temp
				temp.Next = n
				break
			}
			cur = cur.Next
		}
		//前面未插入，所以cur必定为nil
		cur.Next = temp
	}
	return fakehead.Next
}
