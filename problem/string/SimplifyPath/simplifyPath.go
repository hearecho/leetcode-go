package SimplifyPath

import (
	"strings"
)

func simplifyPath(path string) string {
	res := ""
	strs := strings.Split(path,"/")
	stack := make([]string,0)
	for i:=0;i<len(strs);i++ {
		cur := strs[i]
		if cur == "" {
			continue
		}
		//当前显示上一级
		if cur == ".." {
			//stack中有元素,则取出
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if cur !="." && len(cur)> 0 {
			stack = append(stack,cur)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	res = strings.Join(stack,"/")
	return "/"+res
}