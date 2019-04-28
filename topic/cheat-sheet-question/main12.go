package main

import "fmt"

type I interface {
	Get() int
	Set(int)
}

type SS struct {
	Age int
}

func (s SS) Get() int {
	return s.Age
}

func (s SS) Set(age int) {
	s.Age = age
}

func f(i I) {
	switch i := i.(type) {
	case SS:
		fmt.Println(i.Age)
	}
}

func ff(i interface{}) {
	switch i := i.(type) {
	case SS:
		fmt.Println(i.Age)
	}
}

func main() {
	ss := SS{10}
	var i I = ss
	if v, ok := i.(SS); ok {
		fmt.Println(v.Age)
	}
	f(i)

	var ie interface{} = ss
	if v, ok := ie.(SS); ok {
		fmt.Println(v.Age)
	}
	ff(ie)
}
