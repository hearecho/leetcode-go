package dailyProblem

import (
	"fmt"
	"testing"
)

func Test_ladderLength(t *testing.T)  {
	fmt.Println(ladderLength("hit","cog",[]string{"hot","dot","dog","lot","log","cog"}))
}

func Test_sortByBits(t *testing.T)  {
	fmt.Println(sortByBits([]int{2,3,5,7,11,13,17,19}))
}