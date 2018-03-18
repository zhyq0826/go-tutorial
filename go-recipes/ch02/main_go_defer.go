package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	return ioutil.ReadAll(f)
}

func main() {
	fmt.Println(ReadFile("main_go_defer.go"))
}
