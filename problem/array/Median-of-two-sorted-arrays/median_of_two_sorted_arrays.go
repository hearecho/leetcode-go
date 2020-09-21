package Median_of_two_sorted_arrays

//O(max(n,m))时间复杂度，不符合题意
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	l := m + n
	//可能是偶数长度，所以需要记录两个值
	left := -1
	right := -1
	aStart := 0
	bStart := 0
	for i := 0; i<= l/2;i++ {
		left = right
		//数组1 不超出空间并且数组b超出空间或者数组1小于数组2当前数据
		if aStart < m && (bStart >= n || nums1[aStart] < nums2[bStart]) {
			right = nums1[aStart]
			aStart++
		} else {
			right = nums2[bStart]
			bStart++
		}
	}
	//判断是否是复数
	if l& 1 == 0 {
		return float64(left + right) / 2
	} else {
		return float64(right)
	}
}

func optimaize_findMedianSortedArrays(nums1 []int, nums2 []int) float64  {
		// 假设 nums1 的⻓度⼩
		if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
		low, high, k, nums1Mid, nums2Mid := 0, len(nums1),
	(len(nums1)+len(nums2)+1)>>1, 0, 0
		for low <= high {
		// nums1: ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2: ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			// nums1 中的分界线 划多了，要向左边移动
		high = nums1Mid - 1
	} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
		// nums1 中的分界线划少了，要向右边移动
		low = nums1Mid + 1
	} else {
		// 找到合适的划分了，需要输出最终结果了
		// 分为奇数偶数 2 种情况
		break
	}
	}
		midLeft, midRight := 0, 0
		if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
		if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
		if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
		return float64(midLeft+midRight) / 2
}

func max(a,b int)  int{
	if a  > b{
		return a
	} else {
		return b
	}
}

func min(a,b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}