package generateParentheses

//利用深度优先搜索的特性将其看作一棵树
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	//n个括号对，也就是说，左括号n个，右括号n个
	//先加入左括号，之后检查是否左括号数量大于右括号，如果大于则添加右括号
	res := &[]string{}
	generate(n,n,res,"")
	return *res
}

func generate(lindex,rindex int,res *[]string,currstr string)  {
	//终止条件
	if lindex == 0 && rindex==0 {
		*res = append(*res,currstr)
		return
	}
	//添加左括号
	if lindex > 0 {
		generate(lindex-1,rindex,res,currstr+"(")
		//这里不同于其他的回溯，是不需要删除最后一位
	}
	//右括号证明此时字符串中存在的左括号比右括号多，所以需要添加左括号
	if rindex > 0 && lindex < rindex {
		generate(lindex,rindex-1,res,currstr+")")
	}
}
