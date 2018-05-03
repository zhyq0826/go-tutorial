package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//defer 延迟函数的调用，是在 return 之后执行，后进先出 lifo
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
