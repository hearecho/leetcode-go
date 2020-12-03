package contest209

import "sort"

func specialArray(nums []int) int {
	l := len(nums)
	//排序
	sort.Ints(nums)
	for i:=0;i<=l;i++ {
		for j,num := range nums {
			if num >= i {
				if l-j == i {
					return i
				}
				break
			}
		}
	}
	return -1
}
