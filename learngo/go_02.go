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
	data := []*field{{"one"}, {"two"}, {"three"}}

	// v 是指针，go 不需要进行转换，直接把这个指针传递进行了，这个传递发生在 go 直接使用时
	for _, v := range data {
		pp := (*field).print
		fmt.Printf("for %#v\n", v)
		go pp(v)
	}

	time.Sleep(3 * time.Second)
	//goroutines print: three, three, three
}
