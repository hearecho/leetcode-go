package Longestmountain

import (
	"fmt"
	"testing"
)

type param struct {
	A []int
}

func Test_longestMountain(t *testing.T)  {
	params := []param{
		{[]int{2,1,4,7,3,2,5}},
	}
	for _,p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n",p,longestMountain(p.A))
	}
} 
