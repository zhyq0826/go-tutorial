package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Printf("%v\n", &p)
	fmt.Printf("%#v\n", p)
	fmt.Println(p.name)
}

func main() {
	data := []field{{"one"}, {"two"}, {"three"}}

	for _, v := range data {
		// fmt.Printf("%T\n", v)
		// fmt.Printf("%v\n", &v)
		go v.print()
	}

	time.Sleep(2 * time.Second)
}
