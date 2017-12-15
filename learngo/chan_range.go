package main

import (
	"fmt"
	// "time"
)

//close 能够告诉 range 停止迭代，如果一个 chan 被 close 了 则停止迭代，
//而且如果一个 chan 被 close 了，则不能再 send，所以 close 要放到 send 完成之后，
//如果不进行 close range 取万数据之后会等待 ch 中有数据可读，相当于 <-ch,但是已经没有 goroutine 还在 send 所以进行出现死锁
func main() {
	ch := make(chan string)

	go func() {
		ch <- "string1"
		ch <- "string2"
		ch <- "string3"
		// close(ch)
	}()

	// go func() {
	//  time.Sleep(time.Second * 3)
	//  close(ch)
	// }()

	for v := range ch {
		fmt.Println(v)
	}

	// time.Sleep(time.Second * 5)

}
