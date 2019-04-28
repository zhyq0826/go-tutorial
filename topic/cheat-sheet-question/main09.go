package main

import (
	"fmt"
)

func main() {
	//数组可以比较
	fmt.Println([...]string{"1"} == [...]string{"1"})
	//slice 不可以比较
	fmt.Println([]string{"1"} == []string{"1"})
}
