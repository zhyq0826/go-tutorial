package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Printf("func %#v\n", p)
	fmt.Printf("func %T\n", p)
	fmt.Printf("func %v\n", p)
	fmt.Println(p.name)
}

func main() {
	data := []field{{"one"}, {"two"}, {"three"}}

	// 在 for 循环中变量 v 是单一变量，go 使用 v.print 执行时必须进行隐士转换，goroutine 的执行极有可能是发生在 for 循环之后，
	// 此时 v 指向的是最后一个 &v 相当于是 {“three”} 的指针
	for _, v := range data {
		pp := (*field).print
		fmt.Printf("for %#v\n", v)
		go pp(&v)
		// go v.print()
	}

	time.Sleep(3 * time.Second)
	//goroutines print: three, three, three
}
