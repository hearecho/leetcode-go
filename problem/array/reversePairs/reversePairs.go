package reversePairs

func reversePairs(nums []int) int {
	temp := make([]int, len(nums))
	// 分治法 再合并阶段统计逆序对数
	var merge_sort func(l, r int) int
	merge_sort = func(l, r int) int {
		// 划分的终止条件
		if l >= r {
			// 此时的数组长度为1 所以不用划分
			return 0
		}
		// 继续进行划分 先找到中点
		m := l + (r-l)/2
		res := merge_sort(l, m) + merge_sort(m+1, r)
		// 左区间
		i, j := l, m+1
		// 暂时将左区间中全部的数据 拷贝到临时的数组中
		for k := l; k <= r; k++ {
			temp[k] = nums[k]
		}
		// 排序的过程
		for k := l; k <= r; k++ {
			// i == m+1 说明这个时候 左边的区间已经排序完成
			// 所以直接将此时右区间的当前元素直接添加到
			if i == m+1 {
				nums[k] = temp[j]
				j++
				// 同理 j==r+1 的情况是因为右区间已经全部合并完成，所以需要添加左区间当前元素
				// 而temp[i] <= temp[j] 则是为了递增排序，所以此时需要向最终结果里面添加右区间的元素
			} else if j == r+1 || temp[i] <= temp[j] {
				nums[k] = temp[i]
				i++
				// 而剩余的else的情况 就是 temp[i] > temp[j],此时可以用作统计逆序对数
				// 其他情况和 归并排序相同 只是在这个分支下增加统计逆序对数的步骤
			} else {
				nums[k] = temp[j]
				j++
				// 这种情况就是 左区间大于右区间的值
				res += m - i + 1
			}
		}
		return res
	}
	return merge_sort(0, len(nums)-1)
}
