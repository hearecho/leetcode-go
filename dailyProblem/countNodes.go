package dailyProblem

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/**
https://leetcode-cn.com/problems/count-complete-tree-nodes/
*/
//简单层次遍历，但是没有用到完全二叉树的性质
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := 0
	for len(queue) != 0 {
		size := len(queue)
		res += size
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]

	}
	return res
}

//利用完全二叉树的性质，一个完全二叉树的一个节点如果没有左子树，则不可能有右子树
//将树分为一个完全二叉树和一个满二叉树
func countNodes_op(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countLevel(root.Left)
	right := countLevel(root.Right)
	//两者相等，则左子树肯定是满二叉树，则可以直接计算
	if left == right {
		return countNodes_op(root.Right) + (1 << left)
	} else {
		//右子树是满二叉树
		return countNodes_op(root.Left) + (1 << right)
	}
}

//计算左子树的深度,算是高度
func countLevel(root *TreeNode) int {
	res := 0
	for root != nil {
		res++
		root = root.Left
	}
	return res
}
