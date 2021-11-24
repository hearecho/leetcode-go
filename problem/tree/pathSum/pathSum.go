package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	// 根节点到叶子节点之间的路径
	// 可以采用在遍历过程中进行判断，一旦选定就可以将当前的路径添加到结果中
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var bk func(node *TreeNode, curSum int, curPath []int)
	bk = func(node *TreeNode, curSum int, curPath []int) {
		// 回溯截止条件
		// 遇到叶子节点，并且和等于 目标值
		if node.Left == nil && node.Right == nil {
			if curSum+node.Val == target {
				temp := make([]int, len(curPath))
				copy(temp, curPath)
				res = append(res, append(temp, node.Val))
			}
			return
		}
		// 剪枝, 遇到当前节点和已经大于目标值的时候直接返回， 还可能是小数
		// 所以不能剪
		// if curSum+node.Val > target {
		// 	return
		// }
		if node.Left != nil {
			bk(node.Left, curSum+node.Val, append(curPath, node.Val))
		}
		if node.Right != nil {
			bk(node.Right, curSum+node.Val, append(curPath, node.Val))
		}
	}
	bk(root, 0, make([]int, 0))
	return res

}

// 这道题目实际上可以在中序遍历过程中直接得出结果
// 中序遍历  中左右刚好满足计算路径和
