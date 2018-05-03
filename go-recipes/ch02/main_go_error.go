package main

import (
	"fmt"
	"os"
)

func main() {
	// go 拥有 err 类型
	_, err := os.Open("/home/ss")
	if err != nil {
		fmt.Errorf("open file ", err.Error())
	}
}
