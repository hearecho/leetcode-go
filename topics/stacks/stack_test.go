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
