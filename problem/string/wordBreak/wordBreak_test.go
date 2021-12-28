package wordBreak

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak2("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
}
