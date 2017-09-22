package main

import (
	"fmt"
	"sync"
	"time"
)

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
