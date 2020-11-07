package AddBinary

import (
	"math"
	"strconv"
)

func addBinary(a string, b string) string {
	s := ""
	var ex uint8
	diff := math.Abs(float64(len(a) - len(b)))
	for diff > 0{
		if len(a) > len(b) {
			b = "0" + b
		} else {
			a = "0" + a
		}
		diff--
	}

	i := len(a) - 1
	for i>=0{
		anum := a[i]-'0'
		bnum := b[i] - '0'
		t := anum + bnum + ex
		ex = t / 2
		t = t%2
		s = strconv.Itoa(int(t)) +s
		i--
	}
	if ex != 0 {
		s = strconv.Itoa(int(ex)) +s
	}
	return s
}

func reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}