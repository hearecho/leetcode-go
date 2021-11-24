package lowestCommonAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 不同于 二叉搜索树 可以直接使用大小来判断 p,q位于哪边的子树
	// 所有节点的值都是唯一的。
	// 算法的功能就是一路将公共节点返回到顶部， 其中包括自己 本身，因为自己的祖先节点也可以是自己
	if root == nil || root == p || root == q {
		// 递归终止条件
		// root 为nil  或者是找到p q节点
		return root
	}
	// 递归调用左子树
	left := lowestCommonAncestor(root.Left, p, q)
	// 递归调用右子树
	right := lowestCommonAncestor(root.Right, p, q)
	if right == nil && left == nil {
		// 没有找到 p q节点则返回null
		return nil
	}
	if left == nil {
		// 此时的情况是 左子树没有找到目标节点 但是右子树找到了
		// 所以此时返回右子树
		return right
	}
	if right == nil {
		// 与上面情况一样
		return left
	}
	return root
}
