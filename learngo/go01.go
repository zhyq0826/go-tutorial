package main

import (
	"fmt"
	"sync"
	"time"
)

//例子使用的是 lock 用来实现对 map 的 原子操作，通过最简单的 lock 实现
func main() {
	m := make(map[int]string)
	m[2] = "first value"
	var lock sync.Mutex
	go func() {
		lock.Lock()
		m[2] = "second value"
	}()
	time.Sleep(1)
	lock.Lock()
	v := m[2]
	lock.Unlock()
	fmt.Println(v)
}
