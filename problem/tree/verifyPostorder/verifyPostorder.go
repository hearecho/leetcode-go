package verifyPostorder

func verifyPostorder(postorder []int) bool {
	// 可以看出 只给了 后序的序列
	// 划分区间 我们可以知道 区间划分之后的最后一个节点就是这个子树根节点的值
	// 然后再一次为区间继续划分
	var part func(l, r int) bool
	part = func(l, r int) bool {
		if l >= r {
			return true
		}
		root := postorder[r]
		// 找到第一个大于 root的值
		// 如果没有大于的证明 只有左子树
		// 所以要从后往前找
		right_index := l
		for postorder[right_index] < root {
			right_index++
		}
		// 这样左子树的区间为 [l: right_index] 右子树区间为[right_index: r]
		for i := l; i < right_index; i++ {
			if postorder[i] > root {
				return false
			}
		}
		for i := right_index; i < r; i++ {
			if postorder[i] < root {
				return false
			}
		}
		return part(l, right_index-1) && part(right_index, r-1)
	}
	return part(0, len(postorder)-1)
}
