package minArray

func minArray(numbers []int) int {
	//旋转数组朝朝最小值
	// 典型的使用二分法 查找
	// nums[mid] > nums[h]说明最小值在 [mid,h]之间
	// nums[mid] < nums[h]说明最小值在 [l,mid]之间
	// 无法判断 mm 在哪个排序数组中，即无法判断旋转点 xx 在 [i, m][i,m] 还是 [m + 1, j][m+1,j] 区间中。解决方案： 执行 j = j - 1j=j−1 缩小判断范围，
	//
	l, h := 0, len(numbers)-1
	for l < h {
		mid := l + (h-l)/2

		if numbers[mid] > numbers[h] {
			l = mid + 1
		} else if numbers[mid] < numbers[h] {
			h = mid
		} else {
			h--
		}
	}
	return numbers[l]
}
