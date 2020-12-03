# 最大二叉树
> 给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
  二叉树的根是数组中的最大元素。
  左子树是通过数组中最大值左边部分构造出的最大二叉树。
  右子树是通过数组中最大值右边部分构造出的最大二叉树。

### 解题思路
> 递归构建即可,每次都要分割数组，只是在逻辑上上分割，之后递归，降低时间复杂度的话，可以使用栈来模拟递归
```go
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
```