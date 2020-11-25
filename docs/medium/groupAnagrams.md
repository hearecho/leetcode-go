# 字母异位词分组
> 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串
>
### 解题思路
> 按照相同规则排序每个字符串，之后存入哈希表中，哈希表使用排序后的str作为key,相同分组作为value
>
```go
func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string,0)
	tempRes := make([][]string,0)
	for _,str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		sstrs := dict[string(bytes)]
		sstrs = append(sstrs,str)
		dict[string(bytes)] = sstrs
	}
	for _,v := range dict {
		tempRes = append(tempRes,v)
	}
	return tempRes
}
```