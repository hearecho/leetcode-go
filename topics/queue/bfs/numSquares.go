package bfs

//https://leetcode-cn.com/leetbook/read/queue-stack/kfgtt/
//最短距离，最短步数，一般都可以使用广度优先搜索来实现

func NumSquares(n int) int {
	res := 0
	squares := createSquares(n)
	queue := make([]int, 0)
	queue = append(queue, n)
	for len(queue) != 0 {
		size := len(queue)
		res++
		for i := 0; i < size; i++ {
			num := queue[i]
			for i := 0; i < len(squares); i++ {
				if num-squares[i] == 0 {
					return res
				}
				queue = append(queue, num-squares[i])
			}
		}
		queue = queue[size:]
	}
	return res
}

func createSquares(n int) []int {
	res := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		res = append(res, i*i)
	}
	return res
}
