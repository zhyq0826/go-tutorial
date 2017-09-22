package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]string)
	m[2] = "first value"
	var mutex sync.Mutex
	// condition
	cv := sync.NewCond(&mutex)
	updateCompleted := false
	go func() {
		cv.L.Lock()
		m[2] = "second value"
		updateCompleted = true
		cv.Signal() //唤醒一个 goroutine
		cv.L.Unlock()
	}()

	cv.L.Lock() //等待锁释放
	for !updateCompleted {
		cv.Wait() //睡眠等待
	}

	v := m[2]
	cv.L.Unlock()
	fmt.Println(v)
}
