package main

import (
	"fmt"
	// user go package blank alias when init package function
	_ "github.com/zhyq0826/go-tutorial/go-recipes/ch01/lib"
)

func main() {
	fmt.Println("use blank package identify for init")
}
