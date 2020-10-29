package SumNumbers

import (
	"fmt"
	"testing"
)

func Test_sumNumbers(t *testing.T) {
	root := &TreeNode{Val: 4}
	n1 := &TreeNode{Val: 9}
	n2 := &TreeNode{Val: 0}
	n3 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 1}
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4
	fmt.Printf("【output】:%v\n", sumNumbers(root))
}
