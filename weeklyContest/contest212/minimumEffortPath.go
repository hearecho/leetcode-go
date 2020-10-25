package contest210

import (
	"math"
)

var moves = [][]int{{0,1},{0,-1},{1,0},{-1,0}}

func minimumEffortPath(heights [][]int) int {
	//创建dp
	used := make([][]bool,len(heights))
	for i:=0;i<len(heights);i++ {
		used[i] = make([]bool,len(heights[i]))
	}
	//行列
	max := 0
	lastMax := math.MaxInt32
	dfs(0,0,&used,&max,heights,&lastMax)
	return lastMax
}
//深度优先搜索
func dfs(i,j int,used *[][]bool,max *int,heights [][]int,res *int)  {
	if i == len(*used)-1 && j == len((*used)[0])-1 {
		*res = *max
		return
	}
	for _,move := range moves {
		ni := i + move[0]
		nj := j + move[1]
		if ni >= len(*used) || nj >= len((*used)[0]) || ni <0 || nj <0 {
			continue
		}
		if ! (*used)[ni][nj] {
			(*used)[ni][nj]= true
			temp := *max
			if Abs(heights[i][j] - heights[ni][nj]) > *max {
				*max = Abs(heights[i][j] - heights[ni][nj])
			}
			//剪枝
			if *max > *res {
				return
			}
			dfs(ni,nj,used,max,heights,res)
			*max = temp
			(*used)[ni][nj]= false
		}
	}
}
func Abs(a int) int  {
	if a < 0 {
		return -a
	}
	return a
}
