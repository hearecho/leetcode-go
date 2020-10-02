package Valid_Parentheses

//使用栈即可
func isValid(s string) bool {
	//模拟栈
	stack := make([]uint8,0)
	 l := len(s)
	 for i:=0;i<l;i++ {
	 	ch := s[i]
		 //ch为({[的时间直接放入
		 if ch == '(' || ch == '{' || ch == '[' {
			 stack = append(stack,ch)
			 continue
		 }
		 //如果ch为右括号，则判断栈顶元素
		 if len(stack) == 0 {
			 return false
		 }
		 if ch == ')' && stack[len(stack)-1] != '(' {
			 return false
		 }
		 if ch == '}' && stack[len(stack)-1] != '{' {
			 return false
		 }
		 if ch == ']' && stack[len(stack)-1] != '[' {
			 return false
		 }
		 //情况匹配则除去栈顶
		 stack = stack[:len(stack)-1]
	 }
	if len(stack) != 0 {
		return false
	}
	return true
}
