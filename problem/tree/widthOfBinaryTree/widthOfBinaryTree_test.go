package widthOfBinaryTree

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	n5 := &TreeNode{Val: 9}
	n4 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 5}
	n2 := &TreeNode{Val: 2, Right: n5}
	n1 := &TreeNode{Val: 3, Left: n3, Right: n4}
	root := &TreeNode{Val: 1, Left: n1, Right: n2}
	fmt.Println(widthOfBinaryTree(root))
}
