package PartitionLabels

func partitionLabels(S string) []int {
	res := []int{}
	//字母从 'a'-'z'所以可以使用数组代替哈希表
	dict := make([]int, 26)
	//保存字符串中每个字母最后出现的位置
	for i := 0; i < len(S); i++ {
		dict[S[i]-'a'] = i
	}
	left := 0
	for left < len(S) {
		right := left
		for i := left; i < len(S) && i <= right; i++ {
			tempRight := dict[S[i]-'a']
			if tempRight > right {
				right = tempRight
			}
		}
		res = append(res, right-left+1)
		left = right + 1
	}
	return res
}
