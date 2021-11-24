package mirrorTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	// 递归 然后从叶子节点开始慢慢反转
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(temp)
	return root
}

// 一切递归都可以使用辅助栈来模拟递归
func mirrorTreeStack(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		//直接交换即可,由于栈的存在我们仍能按照原来的顺序访问后续的左右节点
		node.Left, node.Right = node.Right, node.Left
	}
	return root
}
