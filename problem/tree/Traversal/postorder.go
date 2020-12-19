package Traversal

func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	var prev *TreeNode
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}

func postorderRecursive(root *TreeNode, res *[]int) {
	//递归方式
	if root == nil {
		return
	}
	preorderRecursive(root.Left, res)
	preorderRecursive(root.Right, res)
	*res = append(*res, root.Val)
}
