package findNthDigit

import (
	"strconv"
)

/**
数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

请写一个函数，求任意第n位对应的数字。
*/

func findNthDigitVio(n int) int {
	// 暴力方法会超时
	if n < 10 {
		return n
	}
	str := ""
	k := 0
	for len(str) < n+1 {
		str = str + strconv.Itoa(k)
		k++
	}
	res := 0
	//if len(str) == n+1 {
	//	res, _ = strconv.Atoi(str[n:])
	//} else {
	//	res, _ = strconv.Atoi(str[n : n+1])
	//}
	res, _ = strconv.Atoi(str[n : n+1])
	return res
}

func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count  // 为了判断最后会在哪个位数中
		start *= 10 // 不同位数的起始数字
		digit += 1
		count = 9 * start * digit // 得出在这个位数 有多少位
	}
	num := start + (n-1)/digit
	s := strconv.Itoa(num)
	res := int(s[(n-1)%digit] - '0')
	return res
}
