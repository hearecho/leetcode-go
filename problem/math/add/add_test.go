package add

import (
	"fmt"
	"testing"
)

func TestBitAdd(t *testing.T) {
	for a := 0; a < 2; a++ {
		for b := 0; b < 2; b++ {
			for c := 0; c < 2; c++ {
				fmt.Println(bitAdd(a, b, c))
			}
		}
	}
}
