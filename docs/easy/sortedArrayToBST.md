# [有序数组转二叉平衡树](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/)
> 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
> 本题中，一个高度平衡二叉树是指一个二叉树每个节点的
> 左右两个子树的高度差的绝对值不超过 1。
>
### 解题思路
```go
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
```