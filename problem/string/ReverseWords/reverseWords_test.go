package ReverseWords

import (
	"fmt"
	"testing"
)

type param struct {
	s string
}

func Test_reverseWords(t *testing.T) {
	params := []param{
		{"  Bob    Loves  Alice   "},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】%v\n", p, reverseWords(p.s))
	}
}

func TestNew(t *testing.T) {
	params := []param{
		{"  Bob    Loves  Alice   "},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】%v\n", p, reverseWordsNew(p.s))
	}
}
