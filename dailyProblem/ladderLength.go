package dailyProblem

//暴力搜索，按照广度优先搜索方式暴力解决
func ladderLength(beginWord string, endWord string, wordList []string) int {
	queue := make([]string, 0)
	res := 1
	used := make([]bool, len(wordList)+1)
	used[0] = true
	queue = append(queue, beginWord)
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			word := queue[i]
			//一旦取出单词与endword相同，则返回
			if word == endWord {
				return res
			}
			//找到距离只有一个字母差距的单词，并添加到队列中
			for i := 0; i < len(wordList); i++ {
				if !used[i+1] && connect(word, wordList[i]) {
					used[i+1] = true
					queue = append(queue, wordList[i])
				}
			}
		}
		queue = queue[l:]
		//进行一层广度搜索则加一，初始是第一层，即res为1
		res++
	}
	return 0
}

func connect(s1, s2 string) bool {
	count := 1
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count--
		}
	}
	return count == 0
}
