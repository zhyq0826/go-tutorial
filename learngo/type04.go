package main

import (
	"fmt"
)

type NewMap map[int]string

func (nm NewMap) add(key int, value string) {
	nm[key] = value
}

func main() {
	var p NewMap = make(map[int]string)
	p.add(10, "a")
	fmt.Println(p)
}
