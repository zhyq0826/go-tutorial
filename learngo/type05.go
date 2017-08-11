package main

import (
	"fmt"
)

type I interface {
	Talk()
}

// existing type 是 I，I 是个接口，可以直接继承 I 的方法，II 等同于 I
type II I

type Person struct {
	name string
}

func (p *Person) Speak() {
	fmt.Println("I am a person")
}

func (p *Person) Talk() {
	fmt.Println("I am talking")
}

func (p Person) Say() {
	fmt.Println("I am saying")
}

type Student Person

type PtrStudent *Person

type HighStudent Student

// embeding type inherit， People 有了 Person 的方法相当于实现了接口 I
type People struct {
	Person
}

func main() {
	var p Person
	p.Speak()
	p.Talk()
	var s Student
	s.name = "jone"
	fmt.Println(s.name)
	var hs HighStudent
	hs.name = "High jone"
	fmt.Println(hs.name)
	var i I
	i = &p
	i.Talk()
	var ii II
	ii = &p
	ii.Talk()
	var ptr PtrStudent = &p
	fmt.Println(ptr.name)
	var people People
	people.name = "people"
	people.Speak()
	people.Talk()
	// s.Speak()
}
