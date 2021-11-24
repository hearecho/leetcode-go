package isSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	/**
		1
	   / \
	  2   2
	 / \ / \
	3  4 4  3
	这样才看做是对称的树, 子树不一定需要对称,
	比较的应该是 nodeleft.left 与 noderight.right 以及 nodeleft.right 与noderight.left的值 是否相同
	*/
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val == b.Val {
		return false
	}
	return symmetric(a.Left, b.Right) && symmetric(a.Right, b.Left)
}
