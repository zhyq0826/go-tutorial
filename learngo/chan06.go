package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 3)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		// close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}
