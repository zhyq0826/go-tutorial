package main

import "fmt"

// bufferd channel 直到 channel 写满才发生阻塞，直到有其他 goroutine 读出其中的值，fib 中的 goroutine 执行到 step 发送
//dang dup3 x 获取到 1 的时候，第二次网 a 中赋值，dup3 开始阻塞
//goroutine 的调度策略
//首先执行 main goroutine，遇到 fib 执行 fib，在 fib 中调用 dup3 执行 dup3，在 dup3 中有 goroutine 执行 goroutine，但是被阻塞，然后直接返回到 fib
// fib 中有 goroutine 执行直到遇到 a 输出阻塞，然后返回 fib 到 main， main 输出 for 遇到阻塞，执行 fib 发现阻塞，然后执行 dup3
func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			fmt.Println("dup3 start ==>")
			x := <-in //取完 in 中的值然后阻塞
			a <- x
			b <- x
			c <- x
			fmt.Println("dup3 end <==")
		}
	}()
	return a, b, c
}

func fib() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		fmt.Println("a first ", <-a)
		for {
			fmt.Println("fib start ==>")
			a_value := <-a
			b_value := <-b
			fmt.Println("a value ", a_value)
			fmt.Println("b value ", b_value)
			x <- a_value + b_value //
			fmt.Println("fib end <==")
		}
	}()
	return out
}

func main() {
	out := fib()
	for i := 0; i < 4; i++ {
		fmt.Println("main start ==>")
		fmt.Println(<-out)
		fmt.Println("main end <==")
	}
}
