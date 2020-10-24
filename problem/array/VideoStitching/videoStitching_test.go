package VideoStitching

import (
	"fmt"
	"testing"
)

type param struct {
	clips [][]int
	T     int
}

func Test_videoStitching(t *testing.T) {
	params := []param{
		{[][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10},
		{[][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}, 9},
	}
	for _, p := range params {
		temp := p
		fmt.Printf("【input】:%v\t【output】:%v\n", temp, videoStitching(p.clips, p.T))
	}
}

func Test_sortSlice(t *testing.T) {
	sl := [][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}
	videoStitching(sl, 10)
	fmt.Println(sl)
}
