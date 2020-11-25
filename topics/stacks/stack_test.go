package stacks

import (
	"fmt"
	"testing"
)

func Test_minStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	minStack.Top()
	fmt.Println(minStack.GetMin())
}
func Test_dailyTemp(t *testing.T) {
	fmt.Println(dailyTemperatures_op([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func Test_evalRPN(t *testing.T) {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
