package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var filePath string
	filePath = "/Users/zhaoyongqiang/t.txt"
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0775)
	// f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	fmt.Println(lines)
}
