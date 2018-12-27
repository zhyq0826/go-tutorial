package main

import (
	"fmt"
)

func main() {
	var ch02 chan int
	if ch02 == nil {
		fmt.Println("ch02 is nil")
	}
}

func func02(a chan int, b chan int) {
	if a == b {
		fmt.Println("a == b")
	}
}
