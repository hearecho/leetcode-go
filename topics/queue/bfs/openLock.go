package bfs

import "strconv"

/*
每个字符串长度固定为4，每个字符有两种变换，所以每个字符串的下一位总共有八种，综合就是总共8种
而最短转换路径，最长使用的就是广度优先搜索
*/
func OpenLock(deadends []string, target string) int {
	step := -1
	queue := make([]string, 0)
	visited := make(map[string]bool, 0)
	for i := 0; i < len(deadends); i++ {
		visited[deadends[i]] = true
	}
	if _, ok := visited["0000"]; ok {
		return -1
	}
	queue = append(queue, "0000")
	for len(queue) != 0 {
		size := len(queue)
		//没过一层就要步数加1，刚开始的0000不算在内
		step++
		for i := 0; i < size; i++ {
			cur := queue[i]
			//当前字符串与目标字符串相同则return
			if cur == target {
				return step
			}
			//取出现在的队头字符串
			for j := 0; j < len(cur); j++ {
				//每个字符的变化,之后再将
				changenum, _ := strconv.Atoi(cur[j : j+1])
				nstr1, nstr2 := "", ""
				if changenum == 9 {
					nstr1 = cur[:j] + strconv.Itoa(0) + cur[j+1:]
				} else {
					nstr1 = cur[:j] + strconv.Itoa(changenum+1) + cur[j+1:]
				}
				if changenum == 0 {
					nstr2 = cur[:j] + strconv.Itoa(9) + cur[j+1:]
				} else {
					nstr2 = cur[:j] + strconv.Itoa(changenum-1) + cur[j+1:]
				}
				if _, ok := visited[nstr1]; !ok {
					queue = append(queue, nstr1)
					visited[nstr1] = true
				}
				if _, ok := visited[nstr2]; !ok {
					queue = append(queue, nstr2)
					visited[nstr2] = true
				}
			}
		}
		queue = queue[size:]
	}
	return -1
}
