package main

import (
	"fmt"
)

func panicRecover() {
	defer fmt.Println("Deferred call - 1")
	defer func() {
		fmt.Println("Defferred call - 2")
		if e := recover(); e != nil {
			fmt.Println("recover with: ", e)
		}
	}()

	panic("just panic")
	fmt.Println("never be reached")
}

func main() {
	fmt.Println("start panic")
	panicRecover()
	fmt.Println("Program regains control after panic recovey")
}
