package num_div

import "math"

//双指针
/**
不是让你这样做的，这道题是为了想到各种情况
*/
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	return dividend / divisor
}

func true_divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32
	}
	sign := 1
	a := int64(dividend)
	b := int64(divisor)
	//异号
	if (a>0&&b<0) || (a<0&&b>0) {
		sign = -1
	}
	if a < 0{
		a = -a

	}
	if b < 0 {
		b = -b
	}
	res := div(a,b)
	if sign > 0{
		if res > math.MaxInt32 {
			return math.MaxInt32
		}
		return int(res)
	}
	return int(-res)
}

func div(a, b int64) int64 {
	if a < b {
		return 0
	}
	var count int64
	count = 1
	tb := b
	for 2*tb <= a {
		//最小解翻倍
		count += count
		//当前测试值也翻倍
		tb += tb
	}
	return count + div(a-tb,b)
}
