package dfs

/**
图的深拷贝，所以就是遍历图，将数值存入新的图中
*/
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node, 0)
	var cg func(node *Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return node
		}
		//节点已经被访问了
		if _, ok := visited[node]; ok {
			return visited[node]
		}
		//新的clone节点
		cloneNode := &Node{
			Val:       node.Val,
			Neighbors: []*Node{},
		}
		//复制节点
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, n)
		}
		return cloneNode
	}
	return cg(node)
}
