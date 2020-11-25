# 求根到叶子节点数字之和
> 给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
> 例如，从根到叶子节点路径 `1->2->3` 代表数字 `123`。

### 注意点
>不是值相加，而是用值组成的数字，之后再将每个路径组成的数字相加

### 解题思路
> 深度优先搜索，先找到每条路径，之后再组合为数字相加
> 如果对于递归理解透彻，那么就可以直接在递归中将数字进行相加，而不用多余的存储空间

```go
// 这就是憨批做法
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nums := make([][]int, 0)
	dps(root, make([]int, 0), &nums)
	fmt.Println(nums)
	tempRes := 0
	for _, num := range nums {
		l := len(num) - 1
		for _, i := range num {
			tempRes += i * int(math.Pow(float64(10), float64(l)))
			l--
		}
	}
	return tempRes
}
func dps(root *TreeNode, curArr []int, tempRes *[][]int) {
	curArr = append(curArr, root.Val)
	if root.Left == nil && root.Right == nil {
		b := make([]int, len(curArr))
		copy(b, curArr)
		*tempRes = append(*tempRes, b)
		return
	}
	if root.Left != nil {
		dps(root.Left, curArr, tempRes)
	}
	if root.Right != nil {
		dps(root.Right, curArr, tempRes)
	}
}

//对递归理解透彻的做法
func sumNumbers_optimaize(root *TreeNode) int {
	return smart(root, 0)
}

func smart(root *TreeNode, a int) int {
	if root == nil {
		return 0
	}
	temp := a*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return temp
	}
	return smart(root.Left, temp) + smart(root.Right, temp)
}
```

