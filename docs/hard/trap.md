### 接雨水
> 从 x 轴开始，给出⼀个数组，数组⾥⾯的数字代表从 (0,0) 点开始，宽度为 1 个单位，⾼度为数组元素
> 的值。如果下⾬了，问这样⼀个容器能装多少单位的⽔？

#### 暴力解法
> 从第二个柱子，开始分别扫描左边和右边，遇到比他大的就停下，之后对每个柱子进行判断累加
```go
//暴力方式
func trap(height []int) int {
	tempRes := 0
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
		tempRes += Min(maxRight, maxLeft) - height[i]
	}
	return tempRes
}

