package main

import (
	"fmt"
	"time"
)

//限制 channel 方向发送端
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

//限制 channel 接收
func printer(c <-chan string) {
	for {
		msg := <-c //接收
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong" //发送
	}
}

func main() {
	var c chan string = make(chan string) // chan 只能用 make 声明
	go pinger(c)
	go printer(c)
	go ponger(c)
	var input string
	fmt.Scanln(&input)
}
