package contest213

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T)  {
	arr := []int{91,4,64,78}
	pi := [][]int{{78},{4,64},{91}}
	fmt.Println(canFormArray(arr,pi))
}

func Test_2(t *testing.T)  {
	fmt.Println(countVowelStrings(2))
}

func Test_3(t *testing.T)  {
	h := []int{14,3,19,3}
	b := 17
	l := 0
	fmt.Println(furthestBuilding(h,b,l))
}

func Test_4(t *testing.T)  {

}
