package causaltest

import (
	"fmt"
	"hckvstore/causal"
	"testing"
)

func TestUpper(t *testing.T) {
	vc := map[string]int32{
		"11": 1,
		"12": 0,
		"13": 1,
	}
	c := causal.MakeTestVC(vc)
	vc2 := map[string]int32{
		"11": 1,
		"12": 0,
		"13": 0,
	}
	fmt.Println(c.IsUpper(vc2))
}

func TestMerge(t *testing.T) {
	vc := map[string]int32{
		"11": 1,
		"12": 3,
		"13": 1,
	}
	c := causal.MakeTestVC(vc)
	vc2 := map[string]int32{
		"11": 1,
		"12": 5,
		"13": 6,
	}
	c.MergeVC(vc2)
	c.PrintVC()
}
