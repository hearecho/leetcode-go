package IncreasingBST

import "fmt"

type TreeNode struct {
	Val int
	Left,Right *TreeNode
}
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

//递归实现
func inorderTraversal_recursive(root *TreeNode)  {
	if root == nil {
		return
	}
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


