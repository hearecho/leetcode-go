package dailyProblem

func moveZeroes(nums []int)  {
	//记录数组中非零的个数
	nozeros := 0
	for _,num := range nums {
		if num != 0{
			nozeros++
		}
	}
	//将所有的非0元素前移
	zeroIndex := 0
	for i:=0;i<len(nums);i++ {
		//如果当前数字不为0，则和zeroIndex交换
		if nums[i] != 0 {
			nums[zeroIndex] = nums[i]
			zeroIndex++
		}
	}
	//后面的全变为0
	for i:=nozeros;i<len(nums);i++ {
		nums[i] = 0
	}
}
