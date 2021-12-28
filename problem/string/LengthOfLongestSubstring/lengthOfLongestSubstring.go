package LengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	// 存储字符的索引
	index := make(map[uint8]int)
	res := 0
	start := 0
	for i := 0; i < len(s); i++ {
		if _, ok := index[s[i]]; ok {
			// 如果存在
			// 这里有问题，之前的数据仍然存在
			temp := index[s[i]]
			for j := start; j <= temp; j++ {
				delete(index, s[j])
			}
			start = temp + 1
		}
		index[s[i]] = i
		if (i - start + 1) > res {
			res = i - start + 1
		}
	}
	return res
}

func lengthOfLongestSubstringOp(s string) int {
	// 存储字符的索引
	// 对删除部分进行优化
	// 我们可以不对前面的数字进行删除，我们对其进行更新即可，
	// 如果该值在哈希表中存在 但是其索引大于start 才进行更新start  否则直接更新该字符所对应的索引
	index := make(map[uint8]int)
	res := 0
	start := 0
	for i := 0; i < len(s); i++ {
		if v, ok := index[s[i]]; ok && v >= start {
			temp := index[s[i]]
			start = temp + 1
		}
		index[s[i]] = i
		if (i - start + 1) > res {
			res = i - start + 1
		}
	}
	return res
}
