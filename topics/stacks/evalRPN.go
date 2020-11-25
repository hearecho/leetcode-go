package stacks

import (
	"strconv"
	"strings"
)

//https://leetcode-cn.com/leetbook/read/queue-stack/gomvm/
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	operators := "+-*/"
	for i := 0; i < len(tokens); i++ {
		if !strings.Contains(operators, tokens[i]) {
			//数字
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		} else {
			//operator
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			tempRes := 0
			stack = stack[:len(stack)-2]
			switch tokens[i] {
			case "*":
				tempRes = num2 * num1
			case "/":
				tempRes = num2 / num1
			case "-":
				tempRes = num2 - num1
			case "+":
				tempRes = num2 + num1
			}
			stack = append(stack, tempRes)
		}
	}
	return stack[0]
}
