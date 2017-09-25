package main

import (
	"fmt"
	"time"
)

func timeout(t chan bool) {
	time.Sleep(5000000000)
	t <- true
}

func readString(s chan string) {
	// s <- "string"
}

//timeout 超时设置
func main() {
	t := make(chan bool)
	s := make(chan string)

	go readString(s)
	go timeout(t)

	select {
	case msg := <-s:
		fmt.Println("recevied", msg)
	case <-t:
		fmt.Println("time out")
	}
}
