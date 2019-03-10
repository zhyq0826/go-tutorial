package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	wr := bufio.NewWriterSize(os.Stdout, 38)
	count := 0
	for {
		wr.WriteString(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second * 1)
		fmt.Println("\ncount ", count)
		count++
		if count > 10 {
			break
		}

	}
	wr.Flush()
}
