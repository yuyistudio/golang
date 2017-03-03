package main

import (
	"fmt"
	"sync"
)


func main() {
	// 首先,map是一种hash表

	var m map[string]int
	fmt.Println(m["key"]) // it's ok to read from a nil map

	var m2 map[string]map[string]int // 不推荐的写法
	fmt.Println(m2["key1"]["key2"]) // 读取很方便
	m2 = map[string]map[string]int{}  // 但是写入很麻烦
	if _, ok := m2["key1"]; ! ok {
		m2["key1"] = map[string]int{}
	}
	m2["key1"]["key2"] = 1

	type Key struct { // struct可以作为map的key,比较时逐条比较其中的字段
		key1, key2 string
	}
	m3 := map[Key]int{}
	m3[Key{"key1","key2"}] = 1 // 写入
	fmt.Println(m3[Key{"key1","key2"}]) // 读取

	var counter = struct{ // 支持并发读取
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}
	counter.Lock()
	counter.m["key"] += 1
	counter.Unlock()
	counter.RLock()
	fmt.Println(counter.m["key"])
	counter.RUnlock()
}

