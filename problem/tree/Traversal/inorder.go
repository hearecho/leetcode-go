package Traversal

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	//由于右子树可能为空
	for len(stack) != 0 || root != nil {
		//遍历左子树
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		//取出栈顶节点
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res

}

func inorderRecursive(root *TreeNode, res *[]int) {
	//递归方式
	if root == nil {
		return
	}
	preorderRecursive(root.Left, res)
	*res = append(*res, root.Val)
	preorderRecursive(root.Right, res)
}
