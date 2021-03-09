package dailyProblem

func removeDuplicates(S string) string {
	stack := make([]uint8, 0)
	res := ""
	for i := 0; i < len(S); i++ {
		if len(stack) == 0 || S[i] != stack[len(stack)-1] {
			stack = append(stack, S[i])
			continue
		}
		//栈不为空且当前元素与栈顶相同，则开始移除栈顶元素直到栈顶元素不与当前元素相同相同
		for len(stack) != 0 && S[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
	}

	for i := 0; i < len(stack); i++ {
		res += string(stack[i])
	}
	return res
}
