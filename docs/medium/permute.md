### 全排列
> 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
#### 关键点
1. 数组没有重复元素
2. 返回数组中也不能出现重复的数组

#### 解题思路
> 深度优先搜索，**需要设置标记，标记已经使用过的元素与一般的回溯相比使用了标记数组**，如第十七题电话号码的字母组合，属于一般回溯问题，不需要记住使用了哪个元素。全排列问题是一类典型回溯题目

```go
//深度优先搜索，需要设置标记
func permute(nums []int) [][]int {
	tempRes := make([][]int, 0)
	used := make([]bool, len(nums))
	backtrack(nums, 0, make([]int, 0), &tempRes, &used)
	return tempRes
}

func backtrack(nums []int, index int, curArr []int, tempRes *[][]int, used *[]bool) {
	if len(curArr) == len(nums) {
		//结束
		b := make([]int, len(nums))
		copy(b, curArr)
		*tempRes = append(*tempRes, b)
		return
	}
	//由于数组中没有重复元素，所以只需要防止元素复用便可以了
	for i := 0; i < len(nums); i++ {
		//如果这个元素没有被用到
		if !(*used)[i] {
			(*used)[i] = true
			curArr = append(curArr, nums[i])
			backtrack(nums, i, curArr, tempRes, used)
			curArr = curArr[:len(curArr)-1]
			(*used)[i] = false
		}
	}

}
```