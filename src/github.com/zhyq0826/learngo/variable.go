package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	a1, b1, c1 := 1, 2, 3 //类型推断
	fmt.Println(a1, b1, c1)

	var e = 10 //一行单值 类型推断
	fmt.Println(e)

	var f float64 = 10 //一行单值，什么类型并初始化
	fmt.Println(f)

	var d, g, k = 10, 11, "k" //一行多值，初始化，类型推断
	fmt.Println(d, g, k)

	var d1, g1, k1 float64 = 10, 11, 12 //一行多值，声明类型并初始化
	fmt.Println(d1, g1, k1)
}
