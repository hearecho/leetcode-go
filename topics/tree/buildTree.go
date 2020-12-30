package tree

/**
根据一棵树的中序遍历与后序遍历构造二叉树
后序遍历结果最后一个是根节点
1. 必须要由一个中序遍历的数字和索引的字典 题目中有说明节点值全部不相同
2. 要先找到根节点，之后将中序节点分开，左边为左子树，右边为右子树
*/
func inandpost(inorder []int, postorder []int) *Node {
	indexs := make(map[int]int, 0)
	for i := 0; i < len(inorder); i++ {
		indexs[inorder[i]] = i
	}
	var build func(start, end int) *Node
	build = func(start, end int) *Node {
		if start > end {
			return nil
		}
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &Node{Val: val}
		rootIndex := indexs[val]
		root.Left = build(start, rootIndex-1)
		root.Right = build(rootIndex+1, end)
		return root
	}

	return build(0, len(postorder)-1)
}
