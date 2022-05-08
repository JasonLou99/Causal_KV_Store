package util

import (
	"log"
	"sync"
)

// Debugging
const Debug = true

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.SetFlags(log.Ldate | log.Lmicroseconds)
		log.Printf(format, a...)
	}
	return
}

/*
sync.Map 相关的函数
*/
func Len(vc sync.Map) int {
	count := 0
	vc.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}
func BecomeMap(vc sync.Map) map[string]int32 {
	res := map[string]int32{}
	vc.Range(func(key, value interface{}) bool {
		res[key.(string)] = value.(int32)
		return true
	})
	return res
}
func BecomeSyncMap(argMap map[string]int32) sync.Map {
	var res sync.Map
	for key, value := range argMap {
		res.Store(key, value)
	}
	return res
}
