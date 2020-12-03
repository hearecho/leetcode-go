package GroupAnagrams

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string,0)
	res := make([][]string,0)
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
		res = append(res,v)
	}
	return res
}

