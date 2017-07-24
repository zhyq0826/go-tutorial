package main

import (
	"fmt"
)

func main() {
	defer func() {
		if ok := recover(); ok != nil {
			fmt.Println("recover")
		}
	}()

	panic("error")

}
