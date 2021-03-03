package dailyProblem

func countBits(num int) []int {
	res := make([]int, 0)
	for i := 0; i <= num; i++ {
		res = append(res, countOne(i))
	}
	return res
}

func countOne(x int) int {
	//计算一个数二进制中1的个数
	res := 0
	for x != 0 {
		x &= x - 1
		res++
	}
	return res
}

//通过上述计算 x的二进制中1的个数 可以推出递推表达式
// count[i] = 1 + count[i & i-1]
func countBits_op(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	for i := 1; i <= num; i++ {
		res[i] = 1 + res[i&(i-1)]
	}

	return res
}
