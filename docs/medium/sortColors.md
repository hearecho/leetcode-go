# [颜色分类](https://leetcode-cn.com/problems/sort-colors/?utm_source=LCUS&utm_medium=ip_redirect_q_uns&utm_campaign=transfer2china)
> 给定一个包含红色、白色和蓝色，一共An 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、A1 和 2 分别表示红色、白色和蓝色。
>

### 解题思路
> 利用三路快排的思想,使用三个指标分贝指向zero的结束位置,two的开始位置，以及遍历索引one

```go
func sortColors(nums []int)  {
	l := len(nums)
	//分别表示0,1,2
	zero, one,two := -1,0,l
	for one < l {
		if nums[one] == 0 && zero < one {
            //前面只可能出现1，所以即使交换了也需要one++
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
```