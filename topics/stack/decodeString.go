package stack

import "strconv"

//https://leetcode-cn.com/leetbook/read/queue-stack/gdwjv/
//使用栈即可
func decodeString(s string) string {
	stack := make([]string, 0)
	res := ""
	for i := 0; i < len(s); {
		//数字左括号直接入栈
		if s[i:i+1] == "]" {
			temp := ""
			//出栈,变为字符串之后再入栈，出栈到左括号
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j] == "[" {
					stack = stack[:len(stack)-1]
					break
				}
				temp = stack[j] + temp
				stack = stack[:len(stack)-1]
			}
			//必定是数字
			count, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			s := temp
			for j := 1; j < count; j++ {
				temp = s + temp
			}
			stack = append(stack, temp)
			i++
		} else {
			//入栈
			//因为数字可能是多位数，所以判断是数字的话要判断距离"["的距离
			if s[i] >= '0' && s[i] <= '9' {
				j := 0
				for s[i+j] >= '0' && s[i+j] <= '9' {
					j++
				}
				stack = append(stack, s[i:i+j])
				i = i + j
			} else {
				stack = append(stack, s[i:i+1])
				i++
			}
		}
	}
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}
