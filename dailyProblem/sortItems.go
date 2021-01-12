package dailyProblem

import "math"

//https://leetcode-cn.com/problems/sort-items-by-groups-respecting-dependencies/

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	//项目  分为有项目组接管 和无项目组接管，
	//而对于无项目组结果如果也没有前置项目，可以最后直接加在项目排序列表结尾
	//而有前置项目组的，则加到前置项目所属的那个小组工作中，之后再按照项目的优先进行排序
	rank := make([]int, n)
	//给所有项目的优先级进行排序
	for _, before := range beforeItems {
		for _, i := range before {
			rank[i]++
		}
	}
	//之后进行分组   创建分组  因为还有一个分组是  属于没有前置也没有小组接手
	// m 个小组 分别用0~m-1进行表示
	groups := make(map[int][]int, 0)
	//返回所属的 group
	var find func(x int, visited []bool) int
	find = func(x int, visited []bool) int {
		if len(beforeItems[x]) == 0 {
			return group[x]
		}
		if visited[x] {
			return math.MinInt32
		}
		visited[x] = true
		return find(beforeItems[x][0], visited)
	}
	for i := 0; i < n; i++ {
		g := find(i, make([]bool, n))
		//存在环 ，所以一定没有解决方案
		if g == math.MinInt32 {
			return []int{}
		}
		groups[g] = append(groups[g], i)
	}
	//对每个小组进行拓扑排序  -1  小组除外最后直接接在最后面就可以
	res := make([]int, 0)
	//var topSort func(items []int)
	//topSort = func(items []int) {
	//	for  {
	//
	//	}
	//}
	for k, v := range groups {
		if k != -1 {
			//进行拓扑排序

			//
			res = append(res, v...)
		}
	}
	//拼接小组
	res = append(res, groups[-1]...)
	return res
}

func topSort(graph [][]int, deg, items []int) (orders []int) {
	q := []int{}
	for _, i := range items {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		from := q[0]
		q = q[1:]
		orders = append(orders, from)
		for _, to := range graph[from] {
			deg[to]--
			if deg[to] == 0 {
				q = append(q, to)
			}
		}
	}
	return
}

func sortItems1(n, m int, group []int, beforeItems [][]int) (ans []int) {
	groupItems := make([][]int, m+n) // groupItems[i] 表示第 i 个组负责的项目列表
	for i := range group {
		if group[i] == -1 {
			group[i] = m + i // 给不属于任何组的项目分配一个组
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	groupGraph := make([][]int, m+n) // 组间依赖关系
	groupDegree := make([]int, m+n)
	itemGraph := make([][]int, n) // 组内依赖关系
	itemDegree := make([]int, n)
	for cur, items := range beforeItems {
		curGroupID := group[cur]
		for _, pre := range items {
			preGroupID := group[pre]
			if preGroupID != curGroupID { // 不同组项目，确定组间依赖关系
				groupGraph[preGroupID] = append(groupGraph[preGroupID], curGroupID)
				groupDegree[curGroupID]++
			} else { // 同组项目，确定组内依赖关系
				itemGraph[pre] = append(itemGraph[pre], cur)
				itemDegree[cur]++
			}
		}
	}

	// 组间拓扑序
	items := make([]int, m+n)
	for i := range items {
		items[i] = i
	}
	groupOrders := topSort(groupGraph, groupDegree, items)
	if len(groupOrders) < len(items) {
		return nil
	}

	// 按照组间的拓扑序，依次求得各个组的组内拓扑序，构成答案
	for _, groupID := range groupOrders {
		items := groupItems[groupID]
		orders := topSort(itemGraph, itemDegree, items)
		if len(orders) < len(items) {
			return nil
		}
		ans = append(ans, orders...)
	}
	return
}
