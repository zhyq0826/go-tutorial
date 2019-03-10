package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	r := strings.NewReader("123 456 k789 123")
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// scanWords 按照 space 进行分割
		advance, token, err = bufio.ScanWords(data, atEOF)
		fmt.Printf("data=%q\n", data)
		fmt.Printf("advance=%d\n", advance)
		fmt.Printf("token=%q\n", token)
		fmt.Printf("atEOF=%t\n", atEOF)
		if strings.Trim(string(token), " ") != "" {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}

		return
	}

	scanner := bufio.NewScanner(r)
	scanner.Split(split)

	for scanner.Scan() {
		fmt.Println("scan text=", scanner.Text())
		fmt.Println("=======")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("%s", err)
	}
}
