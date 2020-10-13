package subStringWithAllWords

import (
	"fmt"
	"testing"
)

type param struct {
	s string
	words []string
}

func Test_findSubstring(t *testing.T)  {
	params := []param{
		{"barfoothefoobarman",[]string{"foo","bar"}},
	}
	for _,p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\t",p,findSubstring(p.s,p.words))
	}
}
