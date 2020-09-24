package romanToInt

func romanToInt(s string) int {
	var roman = map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	l := len(s)
	num := 0
	for i := 0; i < len(s); {
		// 4 或者9的情况
		if i+1 < l && s[i] == 'I' && (s[i+1] == 'V' || s[i+1] == 'X') {
			num += roman[s[i+1]] - 1
			i += 2
			continue
		}
		//40 或者90 的情况
		if i+1 < l && s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C') {
			num += roman[s[i+1]] - 10
			i += 2
			continue
		}
		//400 和900的情况
		if i+1 < l && s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M') {
			num += roman[s[i+1]] - 100
			i += 2
			continue
		}
		//其他情况
		num += roman[s[i]]
		i++
	}
	return num
}

func optimaize_romanToInt(s string) int {
	var roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	if s == "" {
		return 0
	}
	num, lastint, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]
		if num < lastint {
			total = total - num
		} else {
			total = total + num
		}
		lastint = num
	}
	return total
}
