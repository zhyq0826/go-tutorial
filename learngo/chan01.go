package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 3)
	go func() {
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("write to c ", i)
		}
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("reading", <-c)
	}
}
