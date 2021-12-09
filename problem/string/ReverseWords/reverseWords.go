package ReverseWords

import "strings"

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

func reverseWordsNew(s string) string {
	// 最简单的方式就是先分割 之后在进行反转
	strs := make([]string, 0)
	for i := 0; i < len(s); {
		for i < len(s) && s[i] == ' ' {
			i++
		}
		start := i
		for i < len(s) && s[i] != ' ' {
			i++
		}
		strs = append(strs, s[start:i])
	}
	res := ""
	for i := len(strs) - 1; i >= 0; i-- {
		res += strs[i] + " "
	}
	return strings.TrimSpace(res)
}
