package SumNumbers

import (
	"fmt"
	"math"
)

// 这就是憨批做法
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nums := make([][]int, 0)
	dps(root, make([]int, 0), &nums)
	fmt.Println(nums)
	res := 0
	for _, num := range nums {
		l := len(num) - 1
		for _, i := range num {
			res += i * int(math.Pow(float64(10), float64(l)))
			l--
		}
	}
	return res
}
func dps(root *TreeNode, curArr []int, res *[][]int) {
	curArr = append(curArr, root.Val)
	if root.Left == nil && root.Right == nil {
		b := make([]int, len(curArr))
		copy(b, curArr)
		*res = append(*res, b)
		return
	}
	if root.Left != nil {
		dps(root.Left, curArr, res)
	}
	if root.Right != nil {
		dps(root.Right, curArr, res)
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
