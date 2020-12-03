package contest213

func furthestBuilding(heights []int, bricks int, ladders int) int {
	diff := make([]int,0)
	for i:=1;i<len(heights);i++ {
		diff = append(diff,heights[i] -heights[i-1])
	}

	res := 0
	backtrack1(diff,bricks,ladders,0,&res)
	return res
}

func backtrack1(diff []int,bricks int, ladders int,curMax int,res *int) {
	if curMax == len(diff) || (ladders ==0 && bricks < diff[curMax])  {
		if curMax > *res {
			*res = curMax
		}
		return
	}
	if diff[curMax] <=0 {
		backtrack1(diff,bricks,ladders,curMax+1,res)
	}
	if bricks >= diff[curMax] && diff[curMax] > 0{
		backtrack1(diff,bricks-diff[curMax],ladders,curMax+1,res)
	}
	if ladders > 0 && diff[curMax] > 0{
		backtrack1(diff,bricks,ladders-1,curMax+1,res)
	}
}
