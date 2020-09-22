package reverse_int

import "math"

//注意区间值
func reverse(x int) int {
	flag := false
	res := 0
	if x < 0 {
		flag = true
		x = -x
	}
	//反转
	for x != 0 {
		res = res*10 + x%10
		x = x/10
		//2147483647
		if res * 10 > math.MaxInt32 && x!=0 {
			return 0
		}
		if res *10 > math.MaxInt32+1 && flag && x!=0{
			return 0
		}
	}
	if flag {
		res = -res
	}
	return res
}

func reverse7(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}
