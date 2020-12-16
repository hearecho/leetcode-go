package dailyProblem

import (
	"fmt"
	"testing"
)

func Test_ladderLength(t *testing.T) {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func Test_sortByBits(t *testing.T) {
	fmt.Println(sortByBits([]int{2, 3, 5, 7, 11, 13, 17, 19}))
}

func Test_countRangeSum(t *testing.T) {
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
}

func Test_swap(t *testing.T) {
	a, b := 1, 2
	fmt.Println(a, b)
}

func Test_nextPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}
func Test_sortArrayByParityII(t *testing.T) {
	nums := []int{4, 1, 2, 1}
	fmt.Println(sortArrayByParityII(nums))
}

func Test_relativeSortArray(t *testing.T) {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}

func Test_canCompleteCircuit(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func Test_moveZeros(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func Test_findMinArrowShots(t *testing.T) {
	points := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	fmt.Println(findMinArrowShots(points))
}

func Test_sortString(t *testing.T) {
	fmt.Println(sortString("aaaabbbbcccc"))
	fmt.Println(sortString("leetcode"))
}

func Test_maximumGap_op(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	fmt.Println(maximumGap_op(nums))
}

func Test_forSum(t *testing.T) {
	A, B, C, D := []int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}

func Test_isPossible(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 6, 8, 9, 9}
	isPossible(nums)
}

func Test_monotoneIncreasingDigits(t *testing.T) {
	fmt.Println(monotoneIncreasingDigits(921))
}
func Test_wordPattern(t *testing.T) {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}
