package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 3
		close(c)
	}()

	fmt.Println(<-c)
}
