package main

import (
	"fmt"
)

func main() {
	const x = 42
	const (
		a = iota
		b = iota //枚举常量
	)

	fmt.Println(x, a, b)
	const (
		a1 = iota
		b1 //隐式使用 iota
	)
	fmt.Println(a1, b1)

	const (
		a2        = 0
		b2 string = "0" //明确指定类型
	)

	fmt.Println(a2, b2)
}
