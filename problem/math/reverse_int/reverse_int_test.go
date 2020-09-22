package reverse_int

import (
	"fmt"
	"math"
	"testing"
)

type param struct {
	x int
}
func Test_reverse(t *testing.T)  {
	ps := []param {
		{321},
		{-123},
		{0},
		{7463847412},
		{9463847412},
	}
	for _,p := range ps {
		fmt.Printf("【input】:%v       【output】:%d\n",p,reverse(p.x))
	}
}

func Test_MaxInt(t *testing.T)  {
	fmt.Println(math.MaxInt32)
}
