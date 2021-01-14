package dailyProblem

//https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/
//每一位上的值  是在前一位的值 加上 2的i次方
//A[0] 到 A[i] 的第 i 个子数组被解释为一个二进制数（从最高有效位到最低有效位）
//由于数组长读超出范围会造成溢出，所以只需要判断余数即可
func prefixesDivBy5(a []int) []bool {
	ans := make([]bool, len(a))
	x := 0
	for i, v := range a {
		x = (x<<1 | v) % 5
		ans[i] = x == 0
	}
	return ans
}
