# [151. 翻转字符串里的单词](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

> 给定一个字符串，逐个翻转字符串中的每个单词。
>
> 说明：
>
> 无空格字符构成一个 单词 。
> 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
> 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

```go

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
```

