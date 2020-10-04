package contest209

/**
如果一棵二叉树满足下述几个条件，则可以称为 奇偶树 ：

二叉树根节点所在层下标为 0 ，根的子节点所在层下标为 1 ，根的孙节点所在层下标为 2 ，依此类推。
偶数下标 层上的所有节点的值都是 奇 整数，从左到右按顺序 严格递增
奇数下标 层上的所有节点的值都是 偶 整数，从左到右按顺序 严格递减
给你二叉树的根节点，如果二叉树为 奇偶树 ，则返回 true ，否则返回 false 。
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func isEvenOddTree(root *TreeNode) bool {

	//层次遍历
	queue := make([]*TreeNode,0)
	queue = append(queue,root)
	flag := false
	for len(queue) != 0{
		l := len(queue)
		prev := 0
		//一个层面的
		for i:=0;i<l;i++ {
			node := queue[i]
			if i == 0 {
				prev = node.Val
			}
			//偶数层
			if !flag  {
				if node.Val %2 == 0 || (i !=0 && node.Val<=prev) {
					return false
				}
			}
			//奇数层
			if flag {
				if node.Val %2 != 0 || (i !=0 && node.Val>=prev) {
					return false
				}
			}
			if node.Left !=nil {
				queue = append(queue,node.Left)
			}
			if node.Right !=nil {
				queue = append(queue,node.Right)
			}
			prev = node.Val
		}
		queue = queue[l:]
		flag = !flag
	}

	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

