package validateBinaryTreeNodes

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	// false 的几种情况
	// 1. 一个节点有两个不同的父节点
	// 2. 互为父节点
	// 3. 分成了两棵二叉树，没有连在一起，即有两个节点没有父节点
	// 4. 形成一个环  祖先节点

	// 根节点就是在两个数组中均为出现的值
	parent := make(map[int]int)
	rootNum := n
	// 记录父节点
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			if _, ok := parent[leftChild[i]]; leftChild[i] != -1 && ok {
				// 该节点已经存在了父节点，所以不构成二叉树
				return false
			}
			parent[leftChild[i]] = i
		}
		if rightChild[i] != -1 {
			if _, ok := parent[rightChild[i]]; rightChild[i] != -1 && ok {
				return false
			}
			parent[rightChild[i]] = i
		}
	}
	// 遍历父节点
	for k, v := range parent {
		if _, ok := parent[v]; ok {
			if parent[v] == k {
				return false
			}
		}
		rootNum--
	}
	return rootNum == 1
}
