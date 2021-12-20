package connect

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connectBfs(root *Node) *Node {
	// 使用广度优先搜索 也就是层次遍历
	if root == nil {
		return root
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		l := len(queue)
		var pre *Node
		for i := 0; i < l; i++ {
			node := queue[i]
			if pre != nil {
				pre.Next = node
			}
			pre = node
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]
	}
	return root
}

func connectDFS(root *Node) *Node {
	// 层次遍历 频繁的进队出队问题
	// 只需要递归遍历所有的节点，并且传入相邻节点
	// 相邻节点只有三种情况
	if root == nil {
		return root
	}
	var dfs func(left, right *Node)
	dfs = func(left, right *Node) {
		if left == nil || left.Next == right {
			return
		}
		left.Next = right
		dfs(left.Left, left.Right)
		dfs(left.Right, right.Left)
		dfs(right.Left, right.Right)
	}
	dfs(root.Left, root.Right)
	return root
}

func connectOpBFS(root *Node) *Node {
	// 一层一层的处理 但是不使用队列
	// 而是使用链表表示下一层
	// 然后每层遍历可以得到下一层的节点
	// cur 当前链表的第一个节点
	cur := root
	for cur != nil {
		fakehead := &Node{}
		last := fakehead
		// 遍历 当前的链表
		for cur != nil {
			// 因为是完全二叉树 其实判断一个也可以
			// 尾插法
			//if cur.Left != nil && cur.Right != nil {
			//	cur.Left.Next = cur.Right
			//	last.Next = cur.Left
			//	last = cur.Right
			//}
			// 更改为适用于所有的二叉树
			if cur.Left != nil {
				last.Next = cur.Left
				last = cur.Left
			}
			if cur.Right != nil {
				last.Next = cur.Right
				last = cur.Right
			}
			cur = cur.Next
		}
		// 更改头节点
		cur = fakehead.Next
	}
	return root
}
