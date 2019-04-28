package main

import (
	"fmt"
)

func main() {
	// nil 不能赋值给 string
	var x string = nil
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}
