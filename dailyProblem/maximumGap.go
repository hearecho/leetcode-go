package dailyProblem

import (
	"math"
	"sort"
)

//https://leetcode-cn.com/problems/maximum-gap/

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	res := -1
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		res = max(res, nums[i]-nums[i-1])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//正确做法应该是使用桶排序
func maximumGap_op(nums []int) int {
	max_, min_ := arrayInfo(nums)
	if len(nums) < 2 || max_ == min_ {
		return 0
	}
	//确定桶的长度以及桶的个数
	bucketLen := max(1, int(math.Ceil(float64((max_-min_)/(len(nums)-1)))))
	bucketNum := int(math.Ceil(float64((max_-min_)/bucketLen))) + 1
	// min_+ bucketLen 区间是左闭右开，所以只需要记住每个桶的开始区间
	bucket := make([][]int, bucketNum)
	for i := 0; i < bucketNum; i++ {
		bucket[i] = make([]int, 0)
	}
	//遍历排序
	for i := 0; i < len(nums); i++ {
		index := int(math.Ceil(float64((nums[i] - min_) / bucketLen)))
		bucket[index] = append(bucket[index], nums[i])
	}
	//遍历桶 获取当前桶的最小值，减去前一个桶的最大值
	preMax := math.MinInt32
	res := math.MinInt32
	for i := 0; i < len(bucket); i++ {
		curMax, curMin := arrayInfo(bucket[i])
		if len(bucket[i]) != 0 && preMax != math.MinInt32 {
			res = max(res, curMin-preMax)
		}
		if len(bucket[i]) != 0 {
			preMax = curMax
		}
	}
	return res
}

func arrayInfo(nums []int) (int, int) {
	max_, min_ := math.MinInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		max_ = max(max_, nums[i])
		min_ = min(min_, nums[i])
	}
	return max_, min_
}
