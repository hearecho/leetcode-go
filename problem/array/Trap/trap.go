package Trap

//暴力方式
func trap(height []int) int {
	res := 0
	for i := 1; i < len(height)-1; i++ {
		maxLeft, maxRight := 0, 0
		//像左扫描，找到左端最大的
		for j := i; j >= 0; j-- {
			maxLeft = Max(height[j], maxLeft)
		}
		//找右边
		for j := i; j < len(height); j++ {
			maxRight = Max(height[j], maxRight)
		}
		res += Min(maxRight, maxLeft) - height[i]
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
