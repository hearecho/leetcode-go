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
