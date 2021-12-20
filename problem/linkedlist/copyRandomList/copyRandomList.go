package copyRandomList

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	// 先直接再后面复制链表，之后再将新节点的Random指向原来指向节点的后一个
	p := head
	for p != nil {
		next := p.Next
		newNode := &Node{Val: p.Val, Next: next}
		p.Next = newNode
		p = next
	}
	// 复制随机指针
	p = head
	for p != nil && p.Next != nil {
		next := p.Next.Next
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = next
	}
	// 拆分  原来的还不能改
	fakeHead := &Node{}
	last := fakeHead
	p = head
	for p != nil && p.Next != nil {
		next := p.Next.Next
		last.Next = p.Next
		last = p.Next
		p.Next = next
		p = next
	}
	return fakeHead.Next
}
