package add

func add(a int, b int) int {
	// 模拟全加器
	// 之后每一位都进行一次计算
	return 0
}

func bitAdd(a int, b int, c int) (int, int) {
	// 一位全加器 返回结果和进位
	sum := a ^ b ^ c
	cout := (a & b) | (b & c) | (a & c)
	return sum, cout
}
