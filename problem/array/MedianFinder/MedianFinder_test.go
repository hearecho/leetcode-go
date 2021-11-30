package MedianFinder

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	obj := Constructor()
	sum := 0
	for i := 0; i < 10; i++ {
		obj.AddNum(i)
		sum += i
		fmt.Println(obj.left, obj.right)
		fmt.Println("此时的平均值是:", obj.FindMedian())
	}
}
