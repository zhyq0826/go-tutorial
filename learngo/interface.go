package main

import (
	"fmt"
)

//1
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

//3
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Set(v int) { p.i = v }

func main() {
	s := S{}
	f(&s) //4

	var i I
	i = &s
	fmt.Println(i.Get())

	if _, ok := i.(*S); ok {
		fmt.Println("i store *S")
	}

	switch t := i.(type) {
	case *S:
		fmt.Println("i store *S", t)
	case *R:
		fmt.Println("i store *R", t)
	}
}
