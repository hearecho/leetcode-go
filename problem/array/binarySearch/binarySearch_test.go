package binarySearch

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T) {
	fmt.Println(searchLeft([]int{1, 2, 2, 4, 4, 6, 7}, 4))
}
