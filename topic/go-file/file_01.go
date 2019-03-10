package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("msg.txt")
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 126)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d = %q", n, buf)
}
