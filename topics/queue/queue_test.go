package queue

import (
	"fmt"
	"github.com/hearecho/leetcode-go/topics/queue/bfs"
	"testing"
)

func Test_numIslands(t *testing.T) {
	grid := [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(bfs.NumIslands(grid))
}

func Test_openClock(t *testing.T) {
	strs := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	fmt.Println(bfs.OpenLock(strs, target))
}

func Test_numSquares(t *testing.T) {
	n := 13
	fmt.Println(bfs.NumSquares(n))
}

func TestMyQueue(t *testing.T) {
	q := Constructor1()
	q.Push(1)
	q.Push(2)
	q.Peek()
	q.Pop()
	q.Empty()
}

func Test_updateMatrix(t *testing.T) {
	grid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	bfs.UpdateMatrix(grid)
}
