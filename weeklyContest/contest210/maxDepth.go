package contest210

func maxDepth(s string) int {
	if s == "" {
		return 0
	}
	stack := []uint8{}
	max := 0
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			stack = append(stack,'(')
		}
		if s[i] == ')' {
			max = Max(max,len(stack))
			stack = stack[:len(stack)-1]
		}
	}
	return max
}

func Max(a,b int) int {
	if a > b{
		return a
	}
	return b
}
