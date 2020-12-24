package romanToInt

func intToRoman(num int) string {
	var roman = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := ""
	for _, n := range nums {
		for i := 0; i < num/n; i++ {
			res += roman[n]
		}
		num %= n
	}
	return res
}
