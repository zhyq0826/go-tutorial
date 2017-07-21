package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
由于闭包，第一个循环中 i 的变量引用的是 for 作用域中 i 的值，如果 for 循环结束 i 的值是固定的 10
*/
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
