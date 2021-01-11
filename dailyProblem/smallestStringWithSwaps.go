package dailyProblem

import "sort"

//https://leetcode-cn.com/problems/smallest-string-with-swaps/solution/jiao-huan-zi-fu-chuan-zhong-de-yuan-su-b-qdn9/

//并查集

//先扫描pairs 将可以相互连接的放在一起作为一个结合 相互连接就是说 两个pairs中有一个索引是相同的
func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	var find func(int) int
	//一层一层得访问父节点 直到找到根节点
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		//找到根节点才返回
		return fa[x]
	}
	//按秩合并  这个是为了路径压缩
	union := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx == fy {
			return
		}
		//交换排序
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		//设置fy得父节点 也就是说fy与fx是连接在一起得 并以fx为父节点
		fa[fy] = fx
	}

	for _, p := range pairs {
		union(p[0], p[1])
	}

	//属于同一个并查集的字符放在一起  后续进行排序
	groups := map[int][]byte{}
	for i := range s {
		f := find(i)
		groups[f] = append(groups[f], s[i])
	}

	for _, bytes := range groups {
		sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	}

	ans := make([]byte, n)
	//按照groups 进行拼接，因为会存在一些不可能进行交换的位置的字符。
	for i := range ans {
		f := find(i)
		ans[i] = groups[f][0]
		groups[f] = groups[f][1:]
	}
	return string(ans)
}
