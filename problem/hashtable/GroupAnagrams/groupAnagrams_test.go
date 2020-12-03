package GroupAnagrams

import (
	"fmt"
	"testing"
)

type param struct {
	strs []string
}

func Test_groupAnagrams(t *testing.T)  {
	params := []param{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bdt"}},
	}
	for _,p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n",p,groupAnagrams(p.strs))
	}
}