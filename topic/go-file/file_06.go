package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("new.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	file.WriteString(time.Now().Local().String())
}
