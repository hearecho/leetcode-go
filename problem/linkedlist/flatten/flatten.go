package flatten

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	// 遇到child 不为nil 则遍历到下一级，之后返回上一级
	// 类似图 可以使用栈 来暂存上一级的节点
	stack := make([]*Node, 0)
	fakehead := &Node{}
	last := fakehead
	for root != nil || len(stack) > 0 {
		if root != nil {
			next := root.Next
			root.Next = nil
			// 双指针节点
			if last != fakehead {
				root.Prev = last
			}
			last.Next = root
			last = root
			if root.Child != nil {
				// 遍历下一层
				temp := root.Child
				root.Child = nil
				stack = append(stack, next)
				root = temp
			} else {
				// 否则仍遍历该层
				root = next
			}
		} else {
			// 将上一层的节点从栈中取出
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return fakehead.Next
}
