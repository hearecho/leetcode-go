package Add_Two_Numbers

import (
	"fmt"
	"leetcode-go/structure"
	"testing"
)

type params struct {
	l1 *structure.ListNode
	l2 *structure.ListNode
}
func Test_addTwoNumbers(t *testing.T)  {
	ps := []params{
		{structure.GenerateLinkedListTail([]int{305}),structure.GenerateLinkedListTail([]int{4,5,6})},
		{structure.GenerateLinkedListTail([]int{1}),structure.GenerateLinkedListTail([]int{9,9,9})},
		{structure.GenerateLinkedListTail([]int{0,4}),structure.GenerateLinkedListTail([]int{8,0,1})},
	}
	for _,p := range ps{
		fmt.Printf("【input】:%v       【output】:%v\n",p,*addTwoNumbers(p.l1,p.l2))
	}
}
