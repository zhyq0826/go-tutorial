package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) sayName() {
	fmt.Println(p.name)
}

type Student struct {
	Person
	name string
}

func main() {
	p := Person{name: "C"}
	p.sayName()

	s1 := Student{name: "Java"}
	s1.sayName() // sayName 的 receiver 仍然是 Person，而此时 Person 的属性 name 是空

	s2 := Student{name: "Java", Person: Person{name: "VB"}}
	s2.sayName()                // Person 已经有了自己的属性
	fmt.Println(s2.name)        //遮蔽 Person 的属性
	fmt.Println(s2.Person.name) // VB
}
