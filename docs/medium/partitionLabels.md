## 划分字母区间
> 字符串 S 由小写字母组成。我们要把这个字符串
> 划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。
>返回一个表示每个字符串片段的长度的列表。

#### 哈希表解法
> dict用于保存每个字符最后出现的位置，由于限定26个字母所以可以使用数组进行代替
> left即每个分割字符串的起点，right表示的是截断字符串的末尾(会随着子串的增长不断更新)
> 内部循环的判断条件是因为，如果第i个字符小于等于right,则可以进入循环，
> 如果此时这个字符的最后出现的位置大于right，则应该更新right。
> 内部循环结束条件，此时的right刚好是子串的最后一个字符所处的位置，所以下一个子串的起始应该是right+1
```go
func partitionLabels(S string) []int {
	tempRes := []int{}
	//字母从 'a'-'z'所以可以使用数组代替哈希表
	dict := make([]int, 26)
	//保存字符串中每个字母最后出现的位置
	for i := 0; i < len(S); i++ {
		dict[S[i]-'a'] = i
	}
	left := 0
	for left < len(S) {
		right := left
		for i := left; i < len(S) && i <= right; i++ {
			tempRight := dict[S[i]-'a']
			if tempRight > right {
				right = tempRight
			}
		}
		tempRes = append(tempRes, right-left+1)
		left = right + 1
	}
	return tempRes
}
```