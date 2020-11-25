package stacks

//https://leetcode-cn.com/leetbook/read/queue-stack/genw3/
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		temp := T[i]
		for j := i + 1; j < len(T); j++ {
			if T[j] > temp {
				res[i] = j - i
				break
			}
		}
	}
	return res
}

func dailyTemperatures_op(T []int) []int {
	res := make([]int, len(T))
	stack := make([]int, 0)
	for i := 0; i < len(T); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		for top := stack[len(stack)-1]; T[top] < T[i]; {
			res[top] = i - top
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			top = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
