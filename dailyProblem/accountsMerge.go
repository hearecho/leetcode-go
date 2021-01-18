package dailyProblem

import "sort"

//最后的按照字典序排序 可以使用转换为 rune数组 进行排序再转换为字符串进行排序
//分组 可以使用并查集,注意排序数组中第一个字符串是名称

func accountsMerge(accounts [][]string) (ans [][]string) {
	//为了将 全部的邮件地址转换为int数组 便于后面
	emailToIndex := map[string]int{}
	// 邮件地址和名称之间的映射关系
	emailToName := map[string]string{}
	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, has := emailToIndex[email]; !has {
				emailToIndex[email] = len(emailToIndex)
				emailToName[email] = name
			}
		}
	}
	//并查集  父亲节点
	parent := make([]int, len(emailToIndex))
	for i := range parent {
		parent[i] = i
	}
	//find函数
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	//合并函数
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}
	//进行对应转化 合并
	for _, account := range accounts {
		firstIndex := emailToIndex[account[1]]
		for _, email := range account[2:] {
			union(emailToIndex[email], firstIndex)
		}
	}

	indexToEmails := map[int][]string{}
	for email, index := range emailToIndex {
		index = find(index)
		indexToEmails[index] = append(indexToEmails[index], email)
	}
	//排序
	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return
}
