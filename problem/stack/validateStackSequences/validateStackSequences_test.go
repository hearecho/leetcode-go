package validateStackSequences

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
