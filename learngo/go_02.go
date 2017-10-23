package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Printf("func p point to value %v\n", p)
	fmt.Printf("func p store value %p\n", p)
}

func main() {
	data := []*field{{"one"}, {"two"}, {"three"}}

	// 在 for 循环中变量 v 是单一变量，go 使用 v.print 执行时必须进行隐士转换，goroutine 的执行极有可能是发生在 for 循环之后
	for _, v := range data {
		pp := (*field).print
		fmt.Printf("for v value %v\n", v)    // {one}
		fmt.Printf("for v address %p\n", &v) // v 的地址，在这里是相同的
		//go 是按值传递，传递的是 v 的拷贝，而 v 是一个指向 filed 的指针，传递的是指针的拷贝，而这个指针在三次都是不一样的
		go pp(v)
	}

	time.Sleep(3 * time.Second)
	//goroutines print: one, two, three
}
