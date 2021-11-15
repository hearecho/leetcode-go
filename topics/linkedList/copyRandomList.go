package linkedList

type Node struct {
	Val int
	Next *Node
	Random *Node
}
func copyRandomList(head *Node) *Node {
	//原地复制，之后再拆分数组。
	// 先在原链表后面直接增加一个相同的节点，其中随机指针先不赋值
	// 删除原来的节点，并在这个过程中，将随机指针的节点指向原来节点的下一个节点
	if head == nil {
		return nil
	}
	p := head
	for p != nil {
		newNode := &Node{Val: p.Val, Next: p.Next}
		next := p.Next
		p.Next = newNode
		p = next
	}
	// 复制随机指针
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}
	// 拆分
	res := head.Next
	last := res
	p = head
	for last.Next != nil {
		p.Next = p.Next.Next
		last.Next = last.Next.Next
		p = p.Next
		last = last.Next
	}
	p.Next = nil
	return res
}

// 第二中简单方式，直接使用哈希表存储新节点和老节点之间的对应关系，然后复制提取即可
func copyRandomList_map(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	p := head
	for p != nil {
		newNode := &Node{Val: p.Val}
		nodeMap[p] = newNode
		p = p.Next
	}
	//next random节点赋值
	p = head
	for p != nil {
		nodeMap[p].Next = nodeMap[p.Next]
		nodeMap[p].Random = nodeMap[p.Random]
	}
	return nodeMap[head]
}

