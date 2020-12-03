# [递增顺序查找树](https://leetcode-cn.com/problems/increasing-order-search-tree/)
> 给你一个树，请你 按中序遍历 重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点

### 二叉树的中序遍历
```go
//递归实现
func inorderTraversal_recursive(root *TreeNode)  {
	inorderTraversal_recursive(root.Left)
	fmt.Println(root.Val)
	inorderTraversal_recursive(root.Right)
}
//迭代实现
func inorderTraversal_iter(root *TreeNode)  {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		//遍历左节点
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		//取出节点
		root = stack[len(stack)-1]
		fmt.Println(root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
}
```

### 解题思路
```go
//就是进行中序遍历，之后修改右孩子节点让他变成一个链表
func increasingBST(root *TreeNode) *TreeNode {
	ans := &TreeNode{}
	pre := ans
	var inorder func(node *TreeNode)
	inorder = func (node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		node.Left = nil
		pre.Right = node
		pre = node
		inorder(node.Right)
	}
	inorder(root)
	return ans.Right
}
```