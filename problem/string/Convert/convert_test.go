package Convert

import (
	"fmt"
	"testing"
)

type param struct {
	s       string
	numRows int
}

func Test_Convert(t *testing.T) {
	params := []param{
		{"LEETCODEISHIRING", 3},
		{"LEETCODEISHIRING", 4},
		{"A", 1},
	}
	for _, p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n", p, convert(p.s, p.numRows))
	}

}
