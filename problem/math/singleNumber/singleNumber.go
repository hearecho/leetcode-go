package singleNumber

func singleNumber(nums []int) int {
	// 解决单个数字其他数字均为m次问题的通用解法
	// 因为 其他数字出现m次则其他数字在该位置的二进制数字之和 肯定是等于m的
	// 最后各个位置的二进制数字之和相加再对m取余则是剩余单数的二进制表示
	ones, twos := 0, 0
	for _, num := range nums {
		ones = ones ^ num & ^twos
		twos = twos ^ num & ^ones
	}
	return ones
}

func singleNumberNormal(nums []int) int {
	counts := make([]int, 32)
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			counts[i] += num & 1
			num >>= 1
		}
	}
	// 将结果提取出来
	res, m := 0, 3
	for i := 31; i >= 0; i-- {
		res <<= 1
		res |= counts[i] % m
	}
	if counts[31]%m != 0 {
		return ^(res ^ 0xffffffff)
	}
	return res
}
