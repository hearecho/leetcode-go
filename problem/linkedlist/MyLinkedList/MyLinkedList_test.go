package MyLinkedList

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(2)
	linkedList.DeleteAtIndex(1)
	linkedList.AddAtHead(2)
	linkedList.AddAtHead(7)
	linkedList.AddAtHead(3)
	linkedList.AddAtHead(2)
	linkedList.AddAtHead(5)
	linkedList.AddAtTail(5)
	fmt.Println(linkedList.Get(5))
	linkedList.DeleteAtIndex(6)
	linkedList.DeleteAtIndex(4)
}
