package ConstructMaximumBinaryTree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	//先找到最大值
	maxIndex := 0
	for i:=0;i<len(nums);i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[maxIndex]}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}

type TreeNode struct {
	Val int
	Left,Right *TreeNode
}