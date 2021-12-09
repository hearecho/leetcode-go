package countDigitOne

func countDigitOne(n int) int {
	digit, res := 1, 0
	high, cur, low := n/10, n%10, 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else if cur > 1 {
			res += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high = high / 10
		digit *= 10
	}
	return res
}
