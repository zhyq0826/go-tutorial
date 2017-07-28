package main

import (
	"fmt"
)

type I interface {
	Talk()
}

type Person struct {
	name string
}

func (p *Person) Speak() {
	fmt.Println("I am a person")
}

func (p *Person) Talk() {
	fmt.Println("I am talking")
}

type Student Person

type HighStudent Student

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
	var people People
	people.name = "people"
	people.Speak()
	people.Talk()
	// s.Speak()
}
