package main

import (
	"fmt"
	"time"
)

func main() {
	p1 := new(int)
	fmt.Printf("p1 --> %#v \n ", p1)
	fmt.Printf("p1 --> %#v \n ", *p1)

	var p2 *int
	i := 0
	p2 = &i
	fmt.Printf("p2 --> %#v \n ", p2)
	fmt.Printf("p2 --> %#v \n ", *p2)

	var s1 []int
	if s1 == nil {
		fmt.Printf("s1 is nil --> %#v \n ", s1)
	}

	s2 := make([]int, 3)
	if s2 == nil {
		fmt.Printf("s2 is nil --> %#v \n ", s2)
	} else {
		fmt.Printf("s2 is not nill --> %#v \n ", s2)
	}

	modifySlice(s2)
	fmt.Printf("s2 is not nill --> %#v \n ", s2)

	var m1 map[int]string
	if m1 == nil {
		fmt.Printf("m1 is nil --> %#v \n ", m1)
	}

	m2 := make(map[int]string)
	if m2 == nil {
		fmt.Printf("m2 is nil --> %#v \n ", m2)
	} else {
		fmt.Printf("m2 is not nill --> %#v \n ", m2)
	}

	modifyMap(m2)
	fmt.Printf("m2 is not nill --> %#v \n ", m2)

	var c1 chan string
	if c1 == nil {
		fmt.Printf("c1 is nil --> %#v \n ", c1)
	}

	c2 := make(chan string)
	if c2 == nil {
		fmt.Printf("c2 is nil --> %#v \n ", c2)
	} else {
		fmt.Printf("c2 is not nill --> %#v \n ", c2)
	}

	modifyChan(c2)
	time.Sleep(5)
	c2value := <-c2
	fmt.Printf("c2 is not nill --> %#v ", c2value)
}

func modifySlice(s []int) {
	s[0] = 1
}

func modifyMap(m map[int]string) {
	m[0] = "string"
}

func modifyChan(c chan string) {
	c <- "string"
}
