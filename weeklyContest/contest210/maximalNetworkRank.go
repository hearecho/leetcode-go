package contest210

import (
	"fmt"
)

func maximalNetworkRank(n int, roads [][]int) int {
	road_num := make([]int,n)
	for _,road := range roads {
		road_num[road[0]]++
	}
	fmt.Println(road_num)
	return 1
}