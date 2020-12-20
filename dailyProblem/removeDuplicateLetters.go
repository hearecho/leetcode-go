package dailyProblem

func removeDuplicateLetters(s string) string {
	stack := make([]uint8, 0)
	res := ""
	index := make([]int, 26)
	used := make([]bool, 26)
	//获取字符的最后出现的位置
	for i := 0; i < len(s); i++ {
		index[s[i]-'a'] = i
	}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			used[s[i]-'a'] = true
		} else if !used[s[i]-'a'] {
			//
			cur := s[i]
			for j := len(stack) - 1; j >= 0; j-- {
				if !(stack[j] > cur && i < index[stack[j]-'a']) {
					break
				}
				used[stack[j]-'a'] = false
				stack = stack[:j]

			}
			stack = append(stack, cur)
			used[cur-'a'] = true
		}
	}
	//取出栈中元素进行拼接
	for i := 0; i < len(stack); i++ {
		res += string(stack[i])
	}
	return res
}
