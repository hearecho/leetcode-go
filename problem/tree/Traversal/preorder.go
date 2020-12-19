package Traversal

func preorder(root *TreeNode) []int {
	//使用栈来代替系统栈
	//刚好和后续相反，只不过在加入栈的时候就进行了访问
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for len(stack) != 0 || root != nil {
		//遍历左子树 包括根节点
		for root != nil {
			stack = append(stack, root)
			res = append(res, root.Val)
			root = root.Left
		}
		//弹出
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}

func preorderRecursive(root *TreeNode, res *[]int) {
	//递归方式
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorderRecursive(root.Left, res)
	preorderRecursive(root.Right, res)
}
