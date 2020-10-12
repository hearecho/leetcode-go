package contest210

func checkPalindromeFormation(a string, b string) bool {
	if IsPalindrome(a) || IsPalindrome(b) {
		return true
	}
	for i:=0; i< len(a);i++ {
		apre := a[:i]
		bpre := b[:i]
		asuf := a[i:]
		bsuf := b[i:]
		if IsPalindrome(apre+bsuf) || IsPalindrome(bpre+asuf) {
			return true
		}
	}
	return false
}

func IsPalindrome(str string) bool  {
	i,j := 0,len(str)-1
	for i<j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}