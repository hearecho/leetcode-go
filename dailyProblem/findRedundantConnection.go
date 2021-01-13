package dailyProblem

//https://leetcode-cn.com/problems/redundant-connection/

//使用并查集的 find函数,并且改造find函数，即找到一个环
func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	//初始化父节点
	for i := range parent {
		parent[i] = i
	}
	//找父节点
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	//合并函数
	union := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}
	//如果合并过程中出现相同
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}
