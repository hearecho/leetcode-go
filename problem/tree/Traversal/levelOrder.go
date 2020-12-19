package Traversal

func levelOrder(root *TreeNode) [][]int {
	//层次遍历
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	queue = append(queue, root)
	for len(queue) != 0 && root != nil {
		temp := make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		res = append(res, temp)
	}
	return res
}
