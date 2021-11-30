package minNumber

import (
	"sort"
	"strconv"
)

type sortNums []int

func (n sortNums) Len() int { return len(n) }
func (n sortNums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n sortNums) Less(i, j int) bool {
	// 转换为字符串
	s1, s2 := strconv.Itoa(n[i]), strconv.Itoa(n[j])
	// 如果两个两个相结合 结合后数字更小的则应该会更小
	temp1, temp2 := s1+s2, s2+s1
	if temp1 > temp2 {
		return false
	} else {
		return true
	}
}

func minNumber(nums []int) string {
	// 最优情况出现的规则
	// 个位数 看作是后面全部是相同的数字 即 3看作 3(333333) 优先级要小于30 31 32 但是优先级大于34 35等等
	// 即位数不同的用个位的值补足后续不同的值
	sort.Sort(sortNums(nums))
	res := ""
	for _, v := range nums {
		res += strconv.Itoa(v)
	}
	return res
}
