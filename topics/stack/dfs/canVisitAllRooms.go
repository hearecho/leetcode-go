package dfs

//https://leetcode-cn.com/leetbook/read/queue-stack/gle1r/
//深度优先搜索，就是图，遍历全部的图
var (
	num int
	vis []bool
)

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	num = 0
	vis = make([]bool, n)
	dfs1(rooms, 0)
	return num == n
}

func dfs1(rooms [][]int, x int) {
	vis[x] = true
	num++
	for _, it := range rooms[x] {
		if !vis[it] {
			dfs1(rooms, it)
		}
	}
}
