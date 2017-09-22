package main

import (
	"fmt"
)

func main() {
	s := "abcdef"
	for _, v := range s {
		fmt.Println(v)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
