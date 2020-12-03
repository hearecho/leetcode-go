package ReverseWords

func reverseWords(s string) string {
	//如果使用额外空间
	words := make([]string, 0)
	for i := 0; i < len(s); {
		if s[i:i+1] == " " {
			i++
			continue
		} else {
			temp := ""
			for i < len(s) {
				if s[i:i+1] == " " {
					break
				}
				temp += s[i : i+1]
				i++
			}
			words = append(words, temp)
		}
	}
	res := ""
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i] + " "
	}
	return res[:len(res)-1]
}
