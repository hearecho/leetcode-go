package maxSlidingWindow

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func TestDqueue(t *testing.T) {
	fmt.Println(maxSlidingWindowQueue([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
