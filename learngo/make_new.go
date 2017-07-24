package main

import (
	"fmt"
)

func main() {
	p1 := new(int)
	fmt.Println("p1 --> ", p1)
	fmt.Println("p1 point to --> ", *p1)

	var p2 *int
	i := 0
	p2 = &i
	fmt.Println("p2 --> ", p1)
	fmt.Println("p2 point to --> ", *p2)

	var s1 []int
	if s1 == nil {
		fmt.Println("s1 is nil --> ", s1)
	}

	s2 := make([]int, 3)
	if s2 == nil {
		fmt.Println("s2 is nil --> ", s2)
	} else {
		fmt.Println("s2 is not nill --> ", s2)
	}

	var m1 map[int]string
	if m1 == nil {
		fmt.Println("m1 is nil --> ", m1)
	}

	m2 := make(map[int]string)
	if m2 == nil {
		fmt.Println("m2 is nil --> ", m2)
	} else {
		fmt.Println("m2 is not nill --> ", m2)
	}
}
