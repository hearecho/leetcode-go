package Permute

//深度优先搜索，需要设置标记
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	backtrack(nums, make([]int, 0), &res, &used)
	return res
}

func backtrack(nums []int, curArr []int, res *[][]int, used *[]bool) {
	if len(curArr) == len(nums) {
		//结束
		b := make([]int, len(nums))
		copy(b, curArr)
		*res = append(*res, b)
		return
	}
	//由于数组中没有重复元素，所以只需要防止元素复用便可以了
	for i := 0; i < len(nums); i++ {
		//如果这个元素没有被用到
		if !(*used)[i] {
			(*used)[i] = true
			curArr = append(curArr, nums[i])
			backtrack(nums, curArr, res, used)
			curArr = curArr[:len(curArr)-1]
			(*used)[i] = false
		}
	}

}
