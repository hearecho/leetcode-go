package Mypow

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	temp := myPow(x, n/2)
	if n%2 == 0 {
		return temp * temp
	}
	return temp * temp * x
}

//迭代版本 不使用递归
func myPowIter(x float64, n int) float64 {
	res := 1.0
	if n == 0 {
		return res
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for n > 0 {
		if n%2 != 0 {
			// 奇数先乘以x
			res *= x
		}
		// 逐步转化为 更大的数的更小的幂运算。
		x *= x
		n /= 2
	}
	return res
}
