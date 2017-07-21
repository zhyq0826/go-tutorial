package main

import (
	"fmt"
	"runtime"
)

// select 类似于 switch case，select 会选择当前第一个 ready 状态的 channel，如果有多个同时准备好，会随机选一个，如果没有则阻塞直到有一个准备好为之。
func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
