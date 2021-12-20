package widthOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Elem struct {
	Node  *TreeNode
	Depth int
	Pos   int
}

func widthOfBinaryTree(root *TreeNode) int {
	// 只有中间为nil的节点才不算 左右两端的均不算在宽度范围内部
	// 层次遍历
	// 要考虑完全二叉树的性质
	queue := make([]Elem, 0)
	cur_depth, left, ans := 0, 0, 0
	queue = append(queue, Elem{Node: root, Depth: 0, Pos: 0})
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			t := queue[i]
			if t.Node != nil {
				queue = append(queue, Elem{t.Node.Left, t.Depth + 1, t.Pos * 2})
				queue = append(queue, Elem{t.Node.Right, t.Depth + 1, t.Pos*2 + 1})
				if cur_depth != t.Depth {
					cur_depth = t.Depth
					left = t.Pos
				}
				ans = max(t.Pos-left+1, ans)
			}
		}
		queue = queue[l:]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
