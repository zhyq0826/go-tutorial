package main

import (
	"fmt"
)

func main() {
	// Declare an unbuffered channel
	counter := make(chan int)
	// Perform send operation by launching new goroutine
	// go func() {
	// }()
	counter <- 10
	fmt.Println(<-counter) // Receive operation from the channel
}
