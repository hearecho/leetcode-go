package PartitionLabels

import (
	"fmt"
	"testing"
)

type param struct {
	s string
}

func Test_PartitionLabels(t *testing.T)  {
	params := []param{
		{"ababcbacadefegdehijhklij"},
	}

	for _,p := range params {
		fmt.Printf("【input】:%v\t【output】:%v\n",p,partitionLabels(p.s))
	}
}