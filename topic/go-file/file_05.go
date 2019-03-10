package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	str := strings.NewReader(strings.Repeat("ab", 10))
	buf := make([]byte, 2)
	reader := bufio.NewReader(str)
	buf, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d = %q", len(buf), buf)
}
