package validateStackSequences

func validateStackSequences(pushed []int, popped []int) bool {
	// 模拟操作
	stack := make([]int, 0)
	i, j := 0, 0
	for i < len(pushed) {
		if len(stack) == 0 || len(stack) != 0 && stack[len(stack)-1] != popped[j] {
			stack = append(stack, pushed[i])
			i++
		}
		for len(stack) != 0 && stack[len(stack)-1] == popped[j] {
			j++
			stack = stack[:len(stack)-1]
		}
	}
	return j == len(popped)
}
