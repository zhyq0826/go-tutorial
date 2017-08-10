package main

import (
	"fmt"
)

func main() {
	var a, b, c = 0, 0, 0
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&c)
	d := a
	fmt.Println(&d)

	var m map[int]string = map[int]string{
		0: "0",
		1: "1",
	}
	mm := m
	fmt.Printf("%p\n", &m)
	fmt.Printf("%p\n", &mm)
	fmt.Println(m)
	fmt.Println(mm)
	changeMap(m)
	fmt.Printf("%p\n", &m)
	fmt.Printf("%p\n", &mm)
	fmt.Println(m)
	fmt.Println(mm)

	var m1 map[int]string
	makeMap(m1)
	fmt.Println(m1 == nil)

	s1 := []int{0, 1, 2, 3, 4}
	fmt.Printf("%p\n", &s1)
	changeSlice(s1)
	fmt.Println(s1)

	c1 := make(chan int)
	fmt.Printf("%p\n", &c1)
	go func() {
		changeChan(c1)
	}()

	fmt.Println(<-c1)

}

func changeMap(m map[int]string) {
	m[2] = "2"
	fmt.Printf("changeMap func %p\n", m)
}

func changeSlice(s []int) {
	fmt.Printf("changeSlice %p\n", &s)
	s = append(s, 5)
}

func changeChan(c chan int) {
	fmt.Printf("changeChan %p\n", &c)
	c <- 0
}

func makeMap(m map[int]string) {
	m = make(map[int]string)
}
