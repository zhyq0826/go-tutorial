package main

import (
	"fmt"
)

type NameAge struct {
	name string
	age  int
}

func (n *NameAge) PlusAge(x int) {
	n.age += x
}

type SecondNameAge NameAge

type ThirdNameAge struct {
	NameAge
}

func main() {
	var p *int //声明一个指针
	fmt.Printf("%v", p)
	println("")          //(nil)
	fmt.Printf("%#v", p) //(*int)(nil)
	println("")
	fmt.Printf("%T", p) //*int
	println("")

	ptr := new(int) //使用 new 声明一个指针
	*ptr = 1
	fmt.Println("ptr", *ptr) //1

	var i int
	p = &i //获取 i 的地址
	fmt.Printf("%v", p)
	println("")
	fmt.Printf("%v", *p) //输出地址指向的值

	a := new(NameAge)
	a.age = 10
	a.name = "hello"
	fmt.Printf("%v\n", a)
	a.PlusAge(10)
	fmt.Printf("%v\n", a)

	b := new(ThirdNameAge)
	b.age = 10
	b.name = "third"
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", b.NameAge)
}
