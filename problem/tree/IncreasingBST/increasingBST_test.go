package IncreasingBST

import "testing"

func Test_increasingBST(t *testing.T)  {
	root := &TreeNode{Val: 5}
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 6}
	n3 := &TreeNode{Val: 2}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 8}
	root.Left,root.Right = n1,n2
	n1.Left,n1.Right = n3,n4
	n2.Right = n5
	root = increasingBST(root)
}