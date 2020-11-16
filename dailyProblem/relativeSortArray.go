package dailyProblem

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	dict := make(map[int]int,0)
	res := make([]int,0)
	temp := make([]int,0)
	for i:=0;i<len(arr2);i++ {
		dict[arr2[i]] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		return dict[arr1[i]] < dict[arr1[j]]
	})
	for i:=0;i<len(arr1);i++ {
		if _,ok := dict[arr1[i]];ok {
			res = append(res, arr1[i])
		} else {
			temp = append(temp,arr1[i])
		}
	}
	sort.Ints(temp)
	res = append(res,temp...)
	return res
}
