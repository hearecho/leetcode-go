package SortedArrayToBST

type TreeNode struct {
	Val int
	Left,Right *TreeNode
}
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	//每次选择数组的中点然后由于数组为有序的所以肯定符合左小右大的特点
	midIndex := len(nums)/2
	root := &TreeNode{Val: nums[midIndex]}
	root.Left = sortedArrayToBST(nums[:midIndex])
	root.Right = sortedArrayToBST(nums[midIndex+1:])
	return root
}