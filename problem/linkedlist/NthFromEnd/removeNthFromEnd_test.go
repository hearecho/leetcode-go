package NthFromEnd

import (
	"fmt"
	"leetcode-go/structure"
	"testing"
)

type param struct {
	nums []int
	n    int
}

func Test_removeNthFromEnd(t *testing.T)  {
	params := []param{
		{[]int{1,2,3,4,5},2},
	}
	for _,p := range params {
		head := structure.GenerateLinkedListTail(p.nums)
		res := removeNthFromEnd(head,p.n)
		fmt.Printf("【input】:%v\n       【output】:\n",p)
		structure.PrinLinkedList(res)
		fmt.Println()
	}
}
