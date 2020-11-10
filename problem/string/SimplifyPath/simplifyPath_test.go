package SimplifyPath

import (
	"fmt"
	"testing"
)

type param struct {
	path string
}

func Test_SimplifyPath(t *testing.T)  {
	params := []param{
		{"/home//foo/"},
		{"/a/../../b/../c//.//"},
		{"/a//b////c/d//././/.."},
	}
	for _,p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n",p,simplifyPath(p.path))
	}
}