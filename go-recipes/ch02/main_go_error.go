package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/home/ss")
	if err != nil {
		fmt.Errorf("open file ", err.Error())
	}
}
