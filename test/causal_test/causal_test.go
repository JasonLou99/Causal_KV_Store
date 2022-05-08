package causaltest

import (
	"fmt"
	"hckvstore/causal"
	"sync"
	"testing"
)

func TestUpper(t *testing.T) {
	var vc sync.Map
	vc.Store("11", int32(1))
	vc.Store("13", int32(1))
	vc.Store("12", int32(1))
	c := causal.MakeTestVC(vc)
	var vc2 sync.Map
	vc2.Store("11", int32(1))
	vc2.Store("13", int32(1))
	vc2.Store("12", int32(0))
	fmt.Println(c.IsUpper(vc2))
	// fmt.Println(util.BecomeMap(vc))
}

func TestMerge(t *testing.T) {
	var vc sync.Map
	vc.Store("11", int32(1))
	vc.Store("13", int32(1))
	vc.Store("12", int32(1))
	c := causal.MakeTestVC(vc)
	var vc2 sync.Map
	vc2.Store("11", int32(1))
	vc2.Store("13", int32(4))
	vc2.Store("15", int32(1))
	c.MergeVC(vc2)
	c.PrintVC()
}
