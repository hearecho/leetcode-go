package contest210

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := make([]bool,0)
	for i := 0;i<len(l);i++ {
		temp := make([]int,len(nums))
		copy(temp,nums)
		ans = append(ans,check(temp,l[i],r[i]))
	}
	return ans
}

func check(nums []int,l,r int) bool {
	checkNums := nums[l:r+1]
	sort.Ints(checkNums)
	diff := checkNums[1] - checkNums[0]
	for i:=2;i<len(checkNums);i++ {
		if checkNums[i]-checkNums[i-1] != diff {
			return false
		}
	}
	return true
}