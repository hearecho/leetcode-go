package isSubStructure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 向下搜索以及向上回溯
	// 空树不是任意一个树的子结构
	// 搜索的起点是每个树的节点开始搜索一次
	// 我们可以进一步的想 将其转换为更小的问题 就是节点是否相同
	// 整个过程的话 我们会先找到和B根节点相同的节点，然后再进行判断是否是子树

	// 特例  两者中有一个为空
	if A == nil || B == nil {
		return false
	}
	var recur func(a *TreeNode, b *TreeNode) bool
	recur = func(a *TreeNode, b *TreeNode) bool {
		// b为空的时候 说明此时b已经不用匹配了，匹配完成
		if b == nil {
			return true
		}
		// a
		if a == nil || a.Val != b.Val {
			return false
		}
		return recur(a.Right, b.Right) && recur(a.Left, b.Left)
	}
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
