package main

import "fmt"

func main() {
	var nums = []int{1,2,5,7,6,5,4,3}
	res := TwoSum(nums,9)
	for i:=0;i<len(res);i++ {
		fmt.Printf("%d\t",res[i])
	}
}
