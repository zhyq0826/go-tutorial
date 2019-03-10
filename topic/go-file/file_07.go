package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("new.txt", os.O_RDWR|os.O_CREATE, 0775)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	file.WriteString(time.Now().Local().String())
}
