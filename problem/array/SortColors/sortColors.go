package SortColors

func sortColors(nums []int)  {
	l := len(nums)
	//分别表示0,1,2
	zero, one,two := -1,0,l
	for one < l {
		if nums[one] == 0 && zero < one {
			zero++
			nums[zero],nums[one] = nums[one],nums[zero]
			one++
		} else if nums[one] == 2 && two > one {
			//这个不需要one++，因为还要在判断一次
			two--
			nums[two],nums[one] = nums[one],nums[two]
		} else {
			one++
		}
	}

}