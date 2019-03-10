package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("msg.txt")
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 126)
	reader := bufio.NewReader(file)
	buf, err = reader.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d = %q", len(buf), buf)
}
