### 全排列Ⅱ
> 给定一个可包含重复数字的序列，返回所有不重复的全排列。

#### 注意点
1. 和46题相比，数组中可能包含重复元素
2. 结果仍然是不重复的全排列
3. 数组并没有说有序
#### 数组中含有重复元素的处理方式
> 将数组排序，之后通过在遍历的过程中，发现这个元素与前一个元素相同，则可以直接跳过
> 这个可以通过画树状图，理解。
> 
![](http://riyugo.com/i/2020/10/23/r5o2j0.jpg)
```go
if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
			continue
}
//不加去重的结果
// [[1 1 2] [1 2 1] [1 1 2] [1 2 1] [2 1 1] [2 1 1]]
// 加了去重的结果
// [[1 1 2] [1 2 1] [2 1 1]]
```
#### 解题思路permute类似
```go
func backtrack(nums []int, used *[]bool, tempRes *[][]int, curArr []int) {
	if len(curArr) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, curArr)
		*tempRes = append(*tempRes, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
			continue
		}
		if !(*used)[i] {
			(*used)[i] = true
			curArr = append(curArr, nums[i])
			backtrack(nums, used, tempRes, curArr)
			curArr = curArr[:len(curArr)-1]
			(*used)[i] = false
		}
	}
}
```