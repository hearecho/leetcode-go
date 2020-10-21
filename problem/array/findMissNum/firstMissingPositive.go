package findMissNum

func firstMissingPositive(nums []int) int {
	l := len(nums)
	//将负数以及0转变为n+1
	for i := 0; i < l; i++ {
		if nums[i] <= 0 {
			nums[i] = l + 1
		}
	}
	//将其他正数位置上的对应的地方的下标表示的数值变为负数
	for i := 0; i < l; i++ {
		num := abs(nums[i])
		if num <= l {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	//遍历找到第一个正数的下标
	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return l + 1
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
