package structure

import "testing"

func Test_LinedkList(t *testing.T) {
	l := GenerateLinkedListTail([]int{1,2,3,4,5})
	PrinLinkedList(l)
}
