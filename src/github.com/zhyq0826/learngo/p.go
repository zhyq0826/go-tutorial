package main

import (
	"fmt"
)

func main() {
	var p *int //声明一个指针
	fmt.Printf("%v", p)
	println("")
	fmt.Printf("%#v", p)
	println("")
	fmt.Printf("%T", p)

	var i int
	p = &i //获取 i 的地址
	fmt.Printf("%v", p)
	println("")
	fmt.Printf("%v", *p) //输出地址指向的值
}
