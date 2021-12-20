package MyLinkedList

type Node struct {
	Val  int
	Next *Node
}
type MyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func Constructor() MyLinkedList {
	node := &Node{}
	return MyLinkedList{Head: node, Tail: node, Length: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if this.Length <= index {
		return -1
	}
	p := this.Head
	for i := 0; i <= index; i++ {
		p = p.Next
	}
	return p.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}
	node.Next = this.Head.Next
	// 添加第一个节点
	if this.Length == 0 {
		this.Tail = node
	}
	this.Head.Next = node
	this.Length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}
	this.Tail.Next = node
	this.Tail = node
	this.Length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Length {
		return
	}
	if index == this.Length {
		this.AddAtTail(val)
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	p := this.Head
	for i := 0; i < index; i++ {
		p = p.Next
	}
	node := &Node{Val: val, Next: p.Next}
	p.Next = node
	this.Length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Length {
		return
	}
	p := this.Head
	for i := 0; i < index; i++ {
		p = p.Next
	}
	next := p.Next.Next
	if next == nil {
		this.Tail = p
	}
	p.Next = next
	this.Length--
}
